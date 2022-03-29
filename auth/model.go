package auth

type RegistrationModel struct {
	GoogleId  *string `db:"google_id"`
	Username  string  `db:"username"`
	Firstname string  `db:"firstname"`
	Lastname  string  `db:"lastname"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	IsEdited  bool    `db:"is_edited"`
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
	IsEdited  bool    `db:"is_edited"`
	Role      string  `db:"role"`
}

type (
	Profile struct {
		ID        int    `json:"id" db:"id"`
		Username  string `json:"user_name" db:"username"`
		Firstname string `json:"first_name" db:"firstname"`
		Lastname  string `json:"last_name" db:"lastname"`
		Email     string `json:"email" db:"email"`
	}

	UpdateProfileParam struct {
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
	}
	UpdateProfileModel struct {
		ID        string `db:"id"`
		Username  string `db:"username"`
		Firstname string `db:"firstname"`
		Lastname  string `db:"lastname"`
		Email     string `db:"email"`
		IsEdited  bool   `db:"is_edited"`
	}
)
