package user_service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helper"
	"final-project/repository/user_repo"
	"net/http"
)

type UserService interface {
	Register(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(userLoginRequest *dto.UserLoginRequest) (*dto.UserLoginResponse, errs.MessageErr)
	Update(userId int, userUpdate *dto.UserUpdateRequest) (*dto.UserUpdateResponse, errs.MessageErr)
	Delete(userId int) (*dto.DeleteResponse, errs.MessageErr)
	Admin(payload *dto.NewUserRequest) (*dto.AdminResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repo.Repository
}

func NewUserService(userRepo user_repo.Repository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helper.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: payload.Password,
		Role: "member",
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	response, err := us.userRepo.CreateNewUser(user)

	if err != nil {
		return nil, err
	}
	response = &dto.NewUserResponse{
		Id:        response.Id,
		FullName:  response.FullName,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
	}

	return response, nil
}

func (us *userService) Update(userId int, userPayload *dto.UserUpdateRequest) (*dto.UserUpdateResponse, errs.MessageErr) {
	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	updateUser, err := us.userRepo.GetUserById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if updateUser.Id != userId {
		return nil, errs.NewNotFoundError("invalid user")
	}

	user := &entity.User{
		Id:       userId,
		FullName: userPayload.FullName,
		Email:    userPayload.Email,
	}

	response, err := us.userRepo.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	response = &dto.UserUpdateResponse{
		Id:        response.Id,
		FullName:  response.FullName,
		Email:     response.Email,
		UpdatedAt: response.UpdatedAt,
	}

	return response, nil
}

func (us *userService) Login(newLoginRequest *dto.UserLoginRequest) (*dto.UserLoginResponse, errs.MessageErr) {
	err := helper.ValidateStruct(newLoginRequest)

	user, err := us.userRepo.GetUserByEmail(newLoginRequest.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(newLoginRequest.Password)

	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.UserLoginResponse{
		Token: token,
	}

	return &response, nil
}

func (us *userService) Delete(userId int) (*dto.DeleteResponse, errs.MessageErr) {
	user, err := us.userRepo.GetUserById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if user.Id != userId {
		return nil, errs.NewNotFoundError("invalid user")
	}

	us.userRepo.DeleteUser(userId)

	response := dto.DeleteResponse{
		Message: "your account has been successfully deleted",
	}

	return &response, nil
}

func (us *userService) Admin(payload *dto.NewUserRequest) (*dto.AdminResponse, errs.MessageErr) {
	err :=  helper.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	admin := &entity.User{
		FullName: "admin",
		Email: "admin@hacktivate.com",
		Password: "admin477",
		Role: "admin",
	}

	err = admin.HashPassword()

	if err != nil {
		return nil, err
	}

	err = us.userRepo.Admin(admin)

	if err != nil {
		return nil, err
	}

	response := dto.AdminResponse{
		Message: "Seeding admin has been successfully",
	}
	return &response, nil
}


