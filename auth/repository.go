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
    users (google_id, username, firstname, lastname, email, password,is_edited, role_id) value (
		:google_id,
        :username,
        :firstname,
        :lastname,
        :email,
        :password,
		:is_edited,
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

func (r Repository) GetUserByIdGoogle(ctx context.Context, id string) (*UserModel, error) {
	query := selectUser + " WHERE google_id = ? LIMIT 1"
	dest := &UserModel{}
	err := r.db.GetContext(ctx, dest, query, id)
	if err != nil {
		return nil, err
	}
	return dest, nil
}
