package auth

type RegistrationModel struct {
	GoogleId  *string `db:"google_id"`
	Username  string  `db:"username"`
	Firstname string  `db:"firstname"`
	Lastname  string  `db:"lastname"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	IsEdited  *bool   `db:"is_edited"`
	RoleId    int     `db:"role_id"`
}

type UserModel struct {
	ID        int     `db:"id"`
	GoogleId  *string `db:"google_id"`
	Username  string  `db:"username"`
	Firstname string  `db:"firstname"`
	Lastname  string  `db:"lastname"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	IsEdited  *bool   `db:"is_edited"`
	Role      string  `db:"role"`
}
