package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	FindAll() (*[]models.SocialMedia, error)
	FindById(id uint) (*models.SocialMedia, error)
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

func (s *socialMediaRepo) FindAll() (*[]models.SocialMedia, error) {
	var socialMedia []models.SocialMedia
	err := s.db.Find(&socialMedia).Error
	return &socialMedia, err
}

func (s *socialMediaRepo) FindById(id uint) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := s.db.First(&socialMedia, id).Error
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
