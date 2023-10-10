package user_repo

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type Repository interface {
	CreateNewUser(userPayLoad *entity.User) (*dto.NewUserResponse, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
	GetUserById(userId int) (*entity.User, errs.MessageErr) 
	UpdateUser(userPayLoad *entity.User) (*dto.UserUpdateResponse, errs.MessageErr)
}
