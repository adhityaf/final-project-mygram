package services

import (
	"final-project/models"
	"final-project/params"
	"final-project/repositories"
	"net/http"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepo
}

func NewPhotoService(photoRepo repositories.PhotoRepo) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepo,
	}
}

func (p *PhotoService) Create(request *params.CreatePhoto) *params.Response {
	modelPhoto := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
	}

	photo, err := p.photoRepo.Create(&modelPhoto)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "Success post a photo",
		Data:    photo,
	}
}

func (p *PhotoService) FindAll() *params.Response {
	photos, err := p.photoRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	if len(*photos) == 0 {
		return &params.Response{
			Status: http.StatusNotFound,
			Error:  "DATA IS EMPTY",
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve all data",
		Data:    photos,
	}
}

func (p *PhotoService) IsPhotoExist(id uint) bool {
	_, err := p.photoRepo.FindById(id)
	return err == nil
}

func (p *PhotoService) Update(request *params.UpdatePhoto) *params.Response {
	photo, err := p.photoRepo.FindByIdAndAuthId(request.ID, request.UserID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't change this data",
		}
	}

	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.PhotoURL = request.PhotoURL

	photo, err = p.photoRepo.Update(photo)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully updated",
		Data:    photo,
	}
}

func (p *PhotoService) Delete(photoId, authId uint) *params.Response {
	photo, err := p.photoRepo.FindByIdAndAuthId(photoId, authId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "UNAUTHORIZED",
			AdditionalInfo: "You can't delete this data",
		}
	}

	_, err = p.photoRepo.Delete(photo)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully deleted",
	}
}
