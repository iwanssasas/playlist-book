package auth

type UserModel struct {
	Username  string `db:"username"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	RoleId    int    `db:"role_id"`
}
