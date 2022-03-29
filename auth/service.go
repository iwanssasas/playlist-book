package auth

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/db"
	"PLAYLISTBOOK/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"context"

	"github.com/bxcodec/faker/v3"
)

type Service struct {
	repository Repository
}

func NewService() Service {
	return Service{
		repository: Repository{
			db: db.New(),
		},
	}
}

func (s Service) Register(ctx context.Context, params RegisterParams) (*string, error) {
	config := config.Get()
	hasing, err := utils.GeneratePassword(params.Password)
	if err != nil {
		return nil, err
	}

	createdRegister := RegistrationModel{
		Username:  params.Username,
		Firstname: params.Firstname,
		Lastname:  params.Lastname,
		Email:     params.Email,
		Password:  string(hasing),
		RoleId:    config.RoleId,
	}
	err = s.repository.Register(ctx, createdRegister)
	if err != nil {
		return nil, err
	}
	done := "RegisterSucceses"
	return &done, nil
}

func (s Service) Login(ctx context.Context, params LoginParams) (*LoginResponse, error) {
	config := config.Get()
	user, err := s.repository.GetUserByIdentity(ctx, params.Identity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	err = utils.ComparePassword(user.Password, params.Password)
	if err != nil {
		return nil, errors.New("password didn't match")
	}
	data := map[string]string{
		"id":        strconv.Itoa(user.ID),
		"username":  user.Username,
		"firstname": user.Firstname,
		"email":     user.Email,
		"role":      user.Role,
	}
	token, err := utils.GenerateToken(config.Secret, config.ExpiredDuration, data)
	if err != nil {
		return nil, err
	}
	result := &LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Firstname + user.Lastname,
		Email:    user.Email,
		Role:     user.Role,
		Token:    *token,
	}
	return result, nil
}

func (s Service) Ping(ctx context.Context) (string, error) {
	user := "Pong"
	return user, nil
}

func (s Service) googleCallback(ctx context.Context, state string, code string) (*Oauth2GoogleResponse, error) {
	if state != oauthStateString {
		return nil, errors.New("invalid oauth state")
	}
	token, err := getGoogleOauth2().Exchange(ctx, code)
	if err != nil {
		return nil, errors.New("code exchange failed")
	}
	response, err := http.Get(googleApis + token.AccessToken)
	if err != nil {
		return nil, errors.New("failed getting user info")
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	var data Oauth2GoogleResponse
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (s Service) registerGoogleAuth(ctx context.Context, data *Oauth2GoogleResponse) (*LoginResponse, error) {
	config := config.Get()

	var user *UserModel

	user, err := s.repository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		hasing, err := utils.GeneratePassword(faker.Password())
		if err != nil {
			return nil, err
		}

		createdRegisterGoogleOauth := RegistrationModel{
			GoogleId:  &data.ID,
			Username:  faker.Username(),
			Firstname: faker.FirstName(),
			Lastname:  faker.LastName(),
			Email:     data.Email,
			Password:  string(hasing),
			IsEdited:  true,
			RoleId:    config.RoleId,
		}
		fmt.Println(createdRegisterGoogleOauth)
		err = s.repository.RegisterGoogle(ctx, createdRegisterGoogleOauth)
		if err != nil {
			return nil, err
		}

		NewUser, err := s.repository.GetUserByEmail(ctx, data.Email)
		if err != nil {
			return nil, err
		}

		user = NewUser
	}

	if user.GoogleId == nil {
		err = s.repository.UpdateGoogleId(ctx, data.ID, data.Email)
		if err != nil {
			return nil, err
		}
	}

	isiData := map[string]string{
		"id":        fmt.Sprint(user.ID),
		"username":  user.Username,
		"firstname": user.Firstname,
		"email":     user.Email,
		"role":      user.Role,
	}

	token, err := utils.GenerateToken(config.Secret, config.ExpiredDuration, isiData)
	if err != nil {
		return nil, err
	}
	result := &LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Firstname + user.Lastname,
		Email:    user.Email,
		Role:     user.Role,
		Token:    *token,
	}
	return result, nil
}

func (s Service) GetUSer(ctx context.Context, id string) (*Profile, error) {
	get, err := s.repository.GetProfil(ctx, id)
	if err != nil {
		return nil, err
	}

	return get, nil
}

func (s Service) EditUser(ctx context.Context, params UpdateProfileParam, id string) (*string, error) {
	name, _ := s.repository.GetUserByUsername(ctx, params.Username)
	if name != nil {
		return nil, errors.New("username in used")
	}
	email, _ := s.repository.GetUserByEmail(ctx, params.Email)
	if email != nil {
		return nil, errors.New("email in used")
	}
	updateProfile := UpdateProfileModel{
		ID:        id,
		Username:  params.Username,
		Firstname: params.Firstname,
		Lastname:  params.Lastname,
		Email:     params.Email,
		IsEdited:  false,
	}
	err := s.repository.UpdateProfileById(ctx, updateProfile)
	if err != nil {
		return nil, err
	}
	done := "YourProfileHasUpdated"
	return &done, nil
}

func (s Service) DeleteUser(id string) error {
	err := s.repository.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil
}
