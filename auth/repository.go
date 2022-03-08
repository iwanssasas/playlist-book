package auth

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

var (
	insertRegister = `
	INSERT INTO
    users (username, firstname, lastname, email, password, role_id) value (
        :username,
        :firstname,
        :lastname,
        :email,
        :password,
        :role_id
    );
	`
	selectUser = `
	SELECT
		u.id,
		u.username,
		u.firstname,
		u.lastname,
		u.email,
		u.password,
		r.name AS role
	FROM
		users u
    JOIN roles r ON u.role_id = r.id
	`
)

func (r Repository) Register(ctx context.Context, params RegistrationModel) error {
	_, err := r.db.NamedExecContext(ctx, insertRegister, params)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetUserByIdentity(ctx context.Context, identity string) (*UserModel, error) {
	query := selectUser + " WHERE username = ? OR email = ? LIMIT 1"
	dest := &UserModel{}
	err := r.db.GetContext(ctx, dest, query, identity, identity)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
