package models

type SocialMedia struct {
	BaseModel
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
}
