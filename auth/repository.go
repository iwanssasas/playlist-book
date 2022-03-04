package auth

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

var (
	insertUser = `
	INSERT INTO
    users (username, firstname, lastname, email, password, role_id) value (
        :username,
        :firstname,
        :lastname,
        :email,
        :password,
        :role_id
    )
`
)

func (r Repository) PostUser(ctx context.Context, params UserModel) error {
	_, err := r.db.NamedExecContext(ctx, insertUser, params)
	if err != nil {
		return err
	}
	return nil
}
