package params

import "time"

type CreateCommentRs struct {
	Id        uint       `json:"id,omitempty"`
	PhotoId   uint       `json:"photo_id,omitempty"`
	UserId    uint       `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Response
}

type GetCommentRs struct {
	Id        uint           `json:"id,omitempty"`
	PhotoId   uint           `json:"photo_id,omitempty"`
	UserId    uint           `json:"user_id,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	User      UserCommentRs  `json:"User,omitempty"`
	Photo     PhotoCommentRs `json:"Photo,omitempty"`
	Response
}

type UserCommentRs struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoCommentRs struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}

type UpdateCommentRs struct {
	Id        uint       `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Caption   string     `json:"caption,omitempty"`
	PhotoUrl  string     `json:"photo_url,omitempty"`
	UserId    uint       `json:"user_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Response
}
