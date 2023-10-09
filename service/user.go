package service

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
	Update(payLoad *entity.User, userUpdate *dto.UserUpdateRequest) (*dto.UserUpdateResponse, errs.MessageErr)
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

func (us *userService) Update(payLoad *entity.User, userUpdate *dto.UserUpdateRequest) (*dto.UserUpdateResponse, errs.MessageErr) {
	err := helper.ValidateStruct(userUpdate)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		FullName: userUpdate.FullName,
		Email: userUpdate.Email,
	}

	updateUser, err := us.userRepo.UpdateUser(payLoad, user)

	if err != nil {
		return nil, err
	}

	response := &dto.UserUpdateResponse{
		Id: updateUser.Id,
		FullName: updateUser.FullName,
		Email: updateUser.Email,
		UpdatedAt: updateUser.UpdatedAt,
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
