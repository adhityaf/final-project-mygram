package services

import (
	"final-project/models"
	"final-project/params"
	"final-project/repositories"
	"net/http"
)

type CommentService struct {
	commentRepo repositories.CommentRepo
}

func NewCommentService(commentRepo repositories.CommentRepo) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (c *CommentService) FindAll() *params.Response {
	comments, err := c.commentRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve comments all data",
		Payload: comments,
	}
}

func(c *CommentService) Create(request params.CreateComment) *params.Response{
	modelComment := models.Comment{
		Message: request.Message,
	}

	comment, err := c.commentRepo.Create(&modelComment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success create comment",
		Payload: comment,
	}

}
