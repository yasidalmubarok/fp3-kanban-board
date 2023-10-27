package handler

import (
	"final-project/handler/category_handler"
	taks_handler "final-project/handler/task_handler"
	"final-project/handler/user_handler"
	"final-project/infrastructure/config"
	"final-project/infrastructure/database"
	"final-project/repository/category_repo/category_pg"
	"final-project/repository/task_repo/task_pg"
	"final-project/repository/user_repo/user_pg"
	"final-project/service/auth_service"
	"final-project/service/category_service"
	"final-project/service/task_service"
	"final-project/service/user_service"

	"github.com/gin-gonic/gin"
)

func StartApp() {

	config.LoadEnv()

	database.InitiliazeDatabase()
	db := database.GetDatabaseInstance()

	//Dependency Injection
	userRepo := user_pg.NewUserPG(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := user_handler.NewUserHandler(userService)

	taskRepo := task_pg.NewTaskRepo(db)
	categoryRepo := category_pg.NewCategoryRepo(db)

	taskService := task_service.NewTaskService(taskRepo, categoryRepo, userRepo)
	categoryService := category_service.NewCategorySevice(categoryRepo, taskRepo)

	categoryHandler := category_handler.NewCategoryHandler(categoryService)
	taskHandler := taks_handler.NewTaskHandler(taskService)

	authService := auth_service.NewAuthService(userRepo, taskRepo, categoryRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.PUT("/update-account", authService.Authentication(), userHandler.Update)
		userRoute.DELETE("/delete-account", authService.Authentication(), userHandler.Delete)
		userRoute.POST("/admin", userHandler.Admin)
	}

	userRoute = route.Group("/categories")
	{
		userRoute.POST("", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.Create)
		userRoute.GET("", authService.Authentication(), categoryHandler.Get)
		userRoute.PATCH("/:categoryId", authService.Authentication(), categoryHandler.Update)
		userRoute.DELETE("/:categoryId", authService.Authentication(), categoryHandler.Delete)
	}

	userRoute = route.Group("/tasks")
	{
		userRoute.POST("", authService.Authentication(), taskHandler.Create)
		userRoute.GET("", authService.Authentication(), taskHandler.Get)
		userRoute.PUT("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.Update)
		userRoute.PATCH("/update-status/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateByStatus)
	}

	route.Run(":" + config.AppConfig().Port)
}
