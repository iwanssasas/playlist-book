package auth

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/db"
	"PLAYLISTBOOK/utils"

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

	createdRegister := UserModel{
		Username:  params.Username,
		Firstname: params.Firstname,
		Lastname:  params.Lastname,
		Email:     params.Email,
		Password:  string(hasing),
		RoleId:    config.RoleId,
	}
	err = s.repository.PostUser(ctx, createdRegister)
	if err != nil {
		return nil, err
	}
	done := "RegisterSucceses"
	return &done, nil
}
