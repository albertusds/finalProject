package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocMed(socmed *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocMed(idToken uint, socmedId uint) error
	GetSocialMedias() (*[]models.SocialMedia, error)
	UpdateSocialMediaById(socmed *models.SocialMedia, idToken uint) error
	GetSocialMediaById(socmedId uint, idToken uint) (*models.SocialMedia, error)
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepo{db: db}
}

func (sr *socialMediaRepo) CreateSocMed(socmed *models.SocialMedia) (*models.SocialMedia, error) {
	err := sr.db.Create(socmed).Error
	return socmed, err
}

func (sr *socialMediaRepo) DeleteSocMed(idToken uint, socmedId uint) error {
	var socmed models.SocialMedia
	return sr.db.Where("user_id=?", idToken).Delete(&socmed, "id=?", socmedId).Error
}

func (sr *socialMediaRepo) GetSocialMedias() (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := sr.db.Find(&socialMedias).Error
	return &socialMedias, err
}

func (sr *socialMediaRepo) UpdateSocialMediaById(socmed *models.SocialMedia, idToken uint) error {
	var socmedDB models.SocialMedia
	err := sr.db.Model(&socmedDB).Where("id=? AND user_id=?", socmed.ID, idToken).Updates(models.SocialMedia{Name: socmed.Name, SocialMediaUrl: socmed.SocialMediaUrl}).Error
	return err
}

func (sr *socialMediaRepo) GetSocialMediaById(socmedId uint, idToken uint) (*models.SocialMedia, error) {
	var socmed models.SocialMedia
	err := sr.db.Where("user_id=?", idToken).First(&socmed, "id=?", socmedId).Error
	return &socmed, err
}
