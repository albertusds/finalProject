package params

import "time"

type SavePhotoRs struct {
	Id        uint       `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Caption   string     `json:"caption,omitempty"`
	PhotoUrl  string     `json:"photo_url,omitempty"`
	UserId    uint       `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Response
}

type UpdatePhotoRs struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetPhotosRs struct {
	Id        uint         `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserId    uint         `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserPhotosRs `json:"User"`
	Response
}

type UserPhotosRs struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
