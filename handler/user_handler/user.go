package user_handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/service/user_service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

// Register implements UserHandler
// Register godoc
// @Summary User register
// @Description User register
// @Tags Users
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "body request for user register"
// @Success 201 {object} dto.NewUserResponse
// @Router /users/register [post]
func (uh *userHandler) Register(ctx *gin.Context) {
	newUserRequest := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Register(newUserRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// Login implements UserHandler.
// Login godoc
// @Summary User login
// @Description User login
// @Tags Users
// @Accept json
// @Produce json
// @Param RequestBody body dto.UserLoginRequest true "body request for user login"
// @Success 200 {object} dto.UserLoginResponse
// @Router /users/login [post]
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

// Updateements UserHandler.
// Update godoc
// @Summary User Update
// @Description User Update
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param RequestBody body dto.UserUpdateRequest true "body request for user login"
// @Success 200 {object} dto.UserUpdateResponse
// @Router /users/update-account [put]
func (uh *userHandler) Update(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(entity.User)

	var userUpdateRequest = &dto.UserUpdateRequest{}

	if !ok {
		errData := errs.NewBadRequest("Failed get user data!!")
		ctx.AbortWithStatusJSON(errData.Status(), errData)
		return
	}
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		ctx.AbortWithStatusJSON(errData.Status(), errData)
		return
	}

	response, err := uh.userService.Update(userData.Id, userUpdateRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete implements UserHandler.
// Delete godoc
// @Summary Delete User
// @Description Delete Users
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.DeleteResponse
// @Router /users/delete-account [delete]
func (uh *userHandler) Delete(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		errData := errs.NewBadRequest("Failed get user data!!")
		ctx.AbortWithStatusJSON(errData.Status(), errData)
		return
	}
	response, err := uh.userService.Delete(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uh *userHandler) Admin(ctx *gin.Context) {
	newUserRequest := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Admin(newUserRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
