package params

import "time"

type RegisterUserRs struct {
	Age      int    `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Response
}

type LoginUserRs struct {
	Token string `json:"token,omitempty"`
	Response
}

type UpdateUserRs struct {
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}
