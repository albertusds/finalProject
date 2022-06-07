package models

type Photo struct {
	BaseModel
	Title    string    `json:"title"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
}
