package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	DeleteComment(idToken uint, commentId uint) error
	GetComment() (*[]models.Comment, error)
	UpdateCommentById(comment *models.Comment, idToken uint) error
	GetCommentById(commentId, idToken uint) (*models.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepository {
	return &commentRepo{db: db}
}

func (c *commentRepo) CreateComment(comment *models.Comment) (*models.Comment, error) {
	err := c.db.Create(comment).Error
	return comment, err
}

func (c *commentRepo) DeleteComment(idToken uint, commentId uint) error {
	var comment models.Comment
	return c.db.Where("user_id=?", idToken).Delete(&comment, "id=?", commentId).Error
}

func (c *commentRepo) GetComment() (*[]models.Comment, error) {
	var comments *[]models.Comment
	err := c.db.Find(&comments).Error
	return comments, err
}

func (c *commentRepo) GetCommentById(commentId, idToken uint) (*models.Comment, error) {
	var comments *models.Comment
	err := c.db.Where("user_id=?", idToken).First(&comments, "id=?", commentId).Error
	return comments, err
}

func (pr *commentRepo) UpdateCommentById(comment *models.Comment, idToken uint) error {
	var commentDB models.Comment
	err := pr.db.Model(&commentDB).Where("id=? AND user_id=?", comment.ID, idToken).Updates(models.Comment{Message: comment.Message}).Error
	return err
}
