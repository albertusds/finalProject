package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserById(user *models.User) error
	GetUserById(userId uint) (*models.User, error)
	DeleteUserById(userId uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) RegisterUser(user *models.User) (*models.User, error) {
	err := u.db.Create(user).Error
	return user, err
}

func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "email=?", email).Error
	return &user, err
}

func (u *userRepo) GetUserById(userId uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "id=?", userId).Error
	return &user, err
}

func (u *userRepo) UpdateUserById(user *models.User) error {
	var userDB models.User
	err := u.db.Model(&userDB).Where("id=?", user.ID).Updates(models.User{Username: user.Username, Email: user.Email}).Error
	return err
}

func (u *userRepo) DeleteUserById(userId uint) error {
	var user models.User
	return u.db.Delete(&user, "id=?", userId).Error
}
