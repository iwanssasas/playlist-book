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
		users (
			username,
			firstname,
			lastname,
			email,
			password,
			is_edited,
			role_id
		)
	VALUES
		(
			:username,
			:firstname,
			:lastname,
			:email,
			:password,
			:is_edited,
			:role_id
		)
	`
	insertRegisterGoogle = `
	INSERT INTO
	users (
		google_id,
		username,
		firstname,
		lastname,
		email,
		password,
		is_edited,
		role_id
	)
VALUES
	(
		:google_id,
		:username,
		:firstname,
		:lastname,
		:email,
		:password,
		:is_edited,
		:role_id
	)
`
	UpdateGoogleId = `
		UPDATE
		users
	SET
		updated_at = NOW(),
		google_id = ?
	WHERE
		email = ?
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

func (r Repository) RegisterGoogle(ctx context.Context, params RegistrationModel) error {
	_, err := r.db.NamedExecContext(ctx, insertRegisterGoogle, params)
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

func (r Repository) GetUserByEmail(ctx context.Context, email string) (*UserModel, error) {
	query := selectUser + " WHERE email = ?"
	dest := UserModel{}
	err := r.db.GetContext(ctx, &dest, query, email)
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

func (r Repository) UpdateGoogleId(ctx context.Context, google_id, email string) error {
	_, err := r.db.ExecContext(ctx, UpdateGoogleId, google_id, email)
	if err != nil {
		return err
	}
	return nil
}
