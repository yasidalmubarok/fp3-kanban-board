package handler

import (
	"final-project/dto"
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
