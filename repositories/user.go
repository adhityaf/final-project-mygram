package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindAll() (*[]models.User, error)
	FindById(id uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) FindAll() (*[]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	return &users, err
}

func (u *userRepo) FindById(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *userRepo) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userRepo) Update(user *models.User) (*models.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *userRepo) Delete(user *models.User) (*models.User, error) {
	err := u.db.Delete(&user).Error
	return user, err
}
