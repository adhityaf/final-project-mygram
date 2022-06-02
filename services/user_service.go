package services

import (
	"final-project/helpers"
	"final-project/params"
	"final-project/repositories"
	"net/http"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) Create(request *params.RegisterUser) *params.Response {
	_, err := u.userRepo.FindByEmail(request.Email)
	if err == nil {
		// if user found
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "Email already registered",
		}
	}

	user, err := u.userRepo.FindByUsername(request.Username)
	if err == nil {
		// if user found
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "Username must be unique",
		}
	}

	// if email not registered
	// create new user
	hashedPass, _ := helpers.HashPassword(request.Password)

	user.Username = request.Username
	user.Email = request.Email
	user.Password = hashedPass
	user.Age = request.Age

	user, err = u.userRepo.Create(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success register user",
		Payload: user,
	}
}

func (u *UserService) FindAll() *params.Response {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve users all data",
		Payload: users,
	}
}

func (u *UserService) Login(request *params.LoginUser) *params.Response {
	user, err := u.userRepo.FindByEmail(request.Email)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "Email not registered yet, please register first",
		}
	}

	isOk := helpers.ComparePassword(user.Password, request.Password)
	if !isOk {
		return &params.Response{
			Status: http.StatusBadRequest,
			Error:  "Invalid email / password",
		}
	}

	token := helpers.GenerateToken(user.ID, uint(user.Age), user.Email, user.Username)

	// if user found
	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success login",
		Payload: user,
		Token:   token,
	}
}

func (u *UserService) Update(request *params.UpdateUser) *params.Response {
	user, err := u.userRepo.FindById(request.ID)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "User Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	user.Username = request.Username
	user.Email = request.Email

	user, err = u.userRepo.Update(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success update user",
		Payload: user,
	}
}

func (u *UserService) Delete(id uint) *params.Response {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "User Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	_, err = u.userRepo.Delete(user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your account has been successfully deleted",
	}
}
