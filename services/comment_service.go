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

	if len(*comments) == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Error:  "DATA IS EMPTY",
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve comments all data",
		Data:    comments,
	}
}

func (c *CommentService) Create(request *params.CreateComment) *params.Response {
	modelComment := models.Comment{
		Message: request.Message,
		PhotoID: uint(request.PhotoID),
		UserID: request.UserID,
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
		Data:    comment,
	}
}

func (c *CommentService) Update(request params.UpdateComment) *params.Response {
	comment, err := c.commentRepo.FindByIdAndAuthId(request.ID, request.UserID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "You can't change this data",
			AdditionalInfo: err.Error(),
		}
	}

	comment.Message = request.Message

	comment, err = c.commentRepo.Update(comment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully updated",
		Data:    comment,
	}
}

func (c *CommentService) Delete(commentId, authId uint) *params.Response {
	comment, err := c.commentRepo.FindByIdAndAuthId(commentId, authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "You can't delete this data",
			AdditionalInfo: err.Error(),
		}
	}

	_, err = c.commentRepo.Delete(comment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully deleted",
	}
}
