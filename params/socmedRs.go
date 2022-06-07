package params

import "time"

type CreateSocMedRs struct {
	Id             uint       `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	SocialMediaUrl string     `json:"social_media_url,omitempty"`
	UserId         uint       `json:"user_id,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	Response
}

type GetSocMedRs struct {
	SocialMedias []GetSocMedItemRs `json:"social_medias,omitempty"`
	Response
}

type GetSocMedItemRs struct {
	Id             uint         `json:"id"`
	Name           string       `json:"name"`
	SocialMediaUrl string       `json:"social_media_url"`
	UserId         uint         `json:"user_id"`
	CreatedAt      *time.Time   `json:"CreatedAt"`
	UpdatedAt      *time.Time   `json:"UpdatedAt"`
	User           UserSocMedRs `json:"User"`
}

type UserSocMedRs struct {
	Id              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type UpdateSocMedRs struct {
	Id             uint       `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	SocialMediaUrl string     `json:"social_media_url,omitempty"`
	UserId         uint       `json:"user_id,omitempty"`
	UpdatedAt      *time.Time `json:"UpdatedAt,omitempty"`
	Response
}
