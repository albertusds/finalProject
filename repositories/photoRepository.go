package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	SavePhoto(photo *models.Photo) (*models.Photo, error)
	UpdatePhotoById(photo *models.Photo, idToken uint) error
	GetPhotos() (*[]models.Photo, error)
	DeletePhoto(idToken uint, photoId uint) error
	GetPhotoById(photoId uint, idToken uint) (*models.Photo, error)
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepository {
	return &photoRepo{db: db}
}

func (pr *photoRepo) SavePhoto(photo *models.Photo) (*models.Photo, error) {
	err := pr.db.Create(photo).Error
	return photo, err
}

func (pr *photoRepo) GetPhotos() (*[]models.Photo, error) {
	var photos []models.Photo
	err := pr.db.Find(&photos).Error
	return &photos, err
}

func (pr *photoRepo) GetPhotoById(photoId uint, idToken uint) (*models.Photo, error) {
	var photo models.Photo
	err := pr.db.Where("user_id=?", idToken).First(&photo, "id=?", photoId).Error
	return &photo, err
}

func (pr *photoRepo) DeletePhoto(idToken uint, photoId uint) error {
	var photo models.Photo
	return pr.db.Where("user_id=?", idToken).Delete(&photo, "id=?", photoId).Error
}

func (pr *photoRepo) UpdatePhotoById(photo *models.Photo, idToken uint) error {
	var photoDB models.Photo
	err := pr.db.Model(&photoDB).Where("id=? AND user_id=?", photo.ID, idToken).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl}).Error
	return err
}
