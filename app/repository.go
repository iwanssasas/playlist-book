package app

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

var (
	insertBook = `
		INSERT INTO
		books (name, description, quantity, author, created_by)
	VALUES
		(
			:name,
			:description,
			:quantity,
			:author,
			:created_by
		);
	`
)

func (r Repository) SubmitBook(ctx context.Context, params BookModel) error {
	_, err := r.db.NamedExecContext(ctx, insertBook, params)
	if err != nil {
		return err
	}
	return nil
}
