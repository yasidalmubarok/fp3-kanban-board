package auth_service

import (
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/user_repo"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	TaskAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo user_repo.Repository
	// taskRepo taskrepository.TaskRepository
}

// , taskRepo task_repo.Repository
func NewAuthService(userRepo user_repo.Repository) AuthService {
	return &authService{
		userRepo: userRepo,
		// taskRepo: taskRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		if userData.Role != "admin" {
			newError := errs.NewUnauthorizedError("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) TaskAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// userData, ok := ctx.MustGet("userData").(*entity.User)
		// if !ok {
		// 	newError := errs.NewBadRequest("Failed to get user data")
		// 	ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
		// 	return
		// }

		// taskID := ctx.Param("taskID")
		// taskIDUint, err := strconv.ParseUint(taskID, 10, 32)
		// if err != nil {
		// 	newError := errs.NewBadRequest("Task id should be an unsigned integer")
		// 	ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
		// 	return
		// }

		// task, err2 := a.taskRepo.GetTaskByID(uint(taskIDUint))
		// if err2 != nil {
		// 	ctx.AbortWithStatusJSON(err2.StatusCode(), err2)
		// 	return
		// }

		// if task.UserID != userData.ID {
		// 	newError := errs.NewUnauthorized("You're not authorized to modify this task")
		// 	ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
		// 	return
		// }

		// ctx.Next()
	}
}
