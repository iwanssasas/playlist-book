package app

import (
	"PLAYLISTBOOK/db"
	"context"
	"strconv"
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

func (s Service) selectAllBook(ctx context.Context) (*SelectBookResponses, error) {
	dest := SelectBookResponses{}
	get, err := s.repository.SelectAllBook(ctx)
	if err != nil {
		return nil, err
	}

	for _, val := range get {
		dest = append(dest, SelectBookResponse{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Quantity:    val.Quantity,
			Author:      val.Author,
			CreatedAt:   val.CreatedAt,
			CreatedBy:   val.CreatedBy,
			UpdateAt:    val.UpdateAt,
			UpdateBy:    val.UpdateBy,
		})

	}
	return &dest, nil
}

func (s Service) DeleteBookById(stringId string) error {
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return err
	}
	err = s.repository.DeleteBookById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateBookById(ctx context.Context, params BookParams, name string, stringId string) (*string, error) {
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return nil, err
	}
	updateBook := UpdateBookModel{
		Name:        params.Name,
		Description: params.Description,
		Quantity:    params.Quantity,
		Author:      params.Author,
		UpdateBy:    name,
		Id:          id,
	}
	err = s.repository.UpdateBookById(ctx, updateBook, id)
	if err != nil {
		return nil, err
	}
	done := "UpdateBookSucceses"
	return &done, nil
}
