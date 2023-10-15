package auth_service

import (
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/category_repo"
	"final-project/repository/task_repo"
	"final-project/repository/user_repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	TaskAuthorization() gin.HandlerFunc
	CategoryAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo user_repo.Repository
	taskRepo task_repo.Repository
	categoryRepo category_repo.Repository
}

// , taskRepo task_repo.Repository
func NewAuthService(userRepo user_repo.Repository, taskRepo task_repo.Repository, categoryRepo category_repo.Repository) AuthService {
	return &authService{
		userRepo: userRepo,
		taskRepo: taskRepo,
		categoryRepo: categoryRepo,
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

		_, err = a.userRepo.GetUserById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		ctx.Set("userData", user)

		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(entity.User)
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
		userData, ok := ctx.MustGet("userData").(entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		taskId, _ := strconv.Atoi(ctx.Param("taskId"))

		task, err := a.taskRepo.GetTaskById(taskId)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if task.UserId != userData.Id {
			newError := errs.NewUnauthorizedError("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) CategoryAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

		task, err := a.taskRepo.GetTaskById(categoryId)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if task.UserId != userData.Id {
			newError := errs.NewUnauthorizedError("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}
