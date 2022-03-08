package app

import (
	"PLAYLISTBOOK/db"
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

func (s Service) SubmitBook(ctx context.Context, params BookParams, name string) (*string, error) {
	createdBook := BookModel{
		Name:        params.Name,
		Description: params.Description,
		Quantity:    params.Quantity,
		Author:      params.Author,
		CreatedBy:   name,
	}
	err := s.repository.SubmitBook(ctx, createdBook)
	if err != nil {
		return nil, err
	}
	done := "InputBookSucceses"
	return &done, nil
}
