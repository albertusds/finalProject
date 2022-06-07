package models

import (
	"finalproject/helpers"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email"`
	Password     string        `gorm:"not null" json:"password"`
	Age          int           `gorm:"not null" json:"age"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo,omitempty"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment,omitempty"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPass(u.Password)
	return err
}
