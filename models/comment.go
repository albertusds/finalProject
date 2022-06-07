package models

type Comment struct {
	BaseModel
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}
