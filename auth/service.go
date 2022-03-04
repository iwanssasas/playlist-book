package auth

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/db"
	"PLAYLISTBOOK/utils"
	"database/sql"
	"errors"

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
	result := &LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Firstname + user.Lastname,
		Email:    user.Username,
		RoleId:   user.RoleId,
	}
	return result, nil
}
