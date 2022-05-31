package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
)

type CommentRepo interface {
	FindAll() (*[]models.Comment, error)
	FindById(id uint) (*models.Comment, error)
	Create(comment *models.Comment) (*models.Comment, error)
	Update(comment *models.Comment) (*models.Comment, error)
	Delete(comment *models.Comment) (*models.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (c *commentRepo) FindAll() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Find(&comments).Error
	return &comments, err
}

func (c *commentRepo) FindById(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := c.db.First(&comment, id).Error
	return &comment, err
}

func (c *commentRepo) Create(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Create(&comment).Error
	return comment, err
}

func (c *commentRepo) Update(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Save(&comment).Error
	return comment, err
}

func (c *commentRepo) Delete(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Delete(&comment).Error
	return comment, err
}
