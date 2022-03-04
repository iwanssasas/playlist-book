package auth

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	RoleId   *int   `json:"role_id"`
}
