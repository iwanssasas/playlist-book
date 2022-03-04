package auth

type RegisterParams struct {
	Username  string `form:"username"`
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
	Email     string `form:"email"`
	Password  string `form:"password"`
}

type LoginParams struct {
	Identity string `form:"identity"`
	Password string `form:"password"`
}
