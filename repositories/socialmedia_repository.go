package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	FindAll(authId uint) (*[]models.SocialMedia, error)
	FindByIdAndAuthId(socialMediaId, authId uint) (*models.SocialMedia, error)
	Create(socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	Update(socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	Delete(socialMedia *models.SocialMedia) (*models.SocialMedia, error)
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) SocialMediaRepo {
	return &socialMediaRepo{
		db: db,
	}
}

func (s *socialMediaRepo) FindAll(authId uint) (*[]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := s.db.Where("user_id=?",authId).Find(&socialMedia).Error
	return &socialMedia, err
}

func (s *socialMediaRepo) FindByIdAndAuthId(socialMediaId, authId uint) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := s.db.Where("user_id=?", authId).First(&socialMedia, socialMediaId).Error
	return &socialMedia, err
}

func (s *socialMediaRepo) Create(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := s.db.Create(&socialMedia).Error
	return socialMedia, err
}

func (s *socialMediaRepo) Update(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := s.db.Save(&socialMedia).Error
	return socialMedia, err
}

func (s *socialMediaRepo) Delete(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	err := s.db.Delete(&socialMedia).Error
	return socialMedia, err
}
