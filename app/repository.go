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
	selectAllBook = `
	SELECT
		id,
		name,
		description,
		quantity,
		author,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM
		books
	ORDER BY 
		name
	LIMIT ? OFFSET ?;
	`
	deleteBook = `
	DELETE FROM books
	`
	updateBook = `
	UPDATE
		books
	SET
		name = :name,
		description = :description,
		quantity = :quantity,
		author = :author,
		updated_by = :updated_by
	WHERE 
		id = :id;
	`
	selectTotalRow = `
    SELECT COUNT(id) FROM books
    `
)

func (r Repository) SubmitBook(ctx context.Context, params BookModel) error {
	_, err := r.db.NamedExecContext(ctx, insertBook, params)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) SelectAllBook(ctx context.Context, offset int, limit int) (ViewBookModels, error) {
	var dest ViewBookModels
	err := r.db.SelectContext(ctx, &dest, selectAllBook, limit, offset)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func (r Repository) GetTotalRow(ctx context.Context) (int, error) {
	var dest int
	err := r.db.GetContext(ctx, &dest, selectTotalRow)
	if err != nil {
		return 0, err
	}
	return dest, nil
}

func (r Repository) DeleteBookById(id int) error {
	_, err := r.db.Exec(deleteBook+" WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateBookById(ctx context.Context, params UpdateBookModel, id int) error {
	_, err := r.db.NamedExecContext(ctx, updateBook, params)
	if err != nil {
		return err
	}
	return nil
}
