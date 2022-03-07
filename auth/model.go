package auth

type RegistrationModel struct {
	Username  string `db:"username"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	RoleId    int    `db:"role_id"`
}

type UserModel struct {
	ID        int    `db:"id"`
	Username  string `db:"username"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	RoleId    int    `db:"role_id"`
}
