package auth

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/db"
	"PLAYLISTBOOK/utils"
	"database/sql"
	"errors"
	"fmt"

	"context"
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
		"id":       fmt.Sprint(user.ID),
		"username": user.Username,
		"email":    user.Email,
		"role":     fmt.Sprint(user.RoleId),
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
		RoleId:   user.RoleId,
		Token:    *token,
	}
	return result, nil
}

func (s Service) Ping(ctx context.Context) (string, error) {
	user := "Pong"
	return user, nil
}
