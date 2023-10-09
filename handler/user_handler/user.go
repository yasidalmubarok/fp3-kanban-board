package user_handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(ctx *gin.Context) {
	newUserRequest := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	newUserRequest.Role = "member"

	result, err := uh.userService.Register(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (uh *userHandler) Login(ctx *gin.Context) {
	userLoginRequest := &dto.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Login(userLoginRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uh *userHandler) Update(ctx *gin.Context) {
	var userUpdateRequest dto.UserUpdateRequest

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		errData := errs.NewBadRequest("Failed get user data!!")
		ctx.JSON(errData.Status(), errData)
		return
	}
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		ctx.JSON(errData.Status(), errData)
		return
	}

	updateUser, err := uh.userService.Update(userData, &userUpdateRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, updateUser)
}
