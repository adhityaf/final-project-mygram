package services

import (
	"final-project/models"
	"final-project/params"
	"final-project/repositories"
	"fmt"
	"net/http"
)

type SocialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepo
}

func NewSocialMediaService(socialMediaRepo repositories.SocialMediaRepo) *SocialMediaService {
	return &SocialMediaService{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *SocialMediaService) Create(request *params.CreateSocialMedia) *params.Response {
	modelSocialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         request.UserID,
	}

	socialMedia, err := s.socialMediaRepo.Create(&modelSocialMedia)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "Success create new social media",
		Payload: socialMedia,
	}
}

func (s *SocialMediaService) FindAll(authId uint) *params.Response {
	socialMedias, err := s.socialMediaRepo.FindAll(authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve all data",
		Payload: socialMedias,
	}
}

func (s *SocialMediaService) Update(request params.UpdateSocialMedia) *params.Response {
	socialMedia, err := s.socialMediaRepo.FindByIdAndAuthId(request.ID, request.UserID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't change this data",
		}
	}

	socialMedia.Name = request.Name
	socialMedia.SocialMediaURL = request.SocialMediaURL

	socialMedia, err = s.socialMediaRepo.Update(socialMedia)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Update data social media with id %d success", request.ID),
		Payload: socialMedia,
	}
}

func (s *SocialMediaService) Delete(socialMediaId, authId uint) *params.Response {
	socialMedia, err := s.socialMediaRepo.FindByIdAndAuthId(socialMediaId, authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't delete this data",
		}
	}

	_, err = s.socialMediaRepo.Delete(socialMedia)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your social media has been successfully deleted",
	}
}
