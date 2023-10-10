package handler

import (
	"final-project/handler/category_handler"
	"final-project/handler/user_handler"
	"final-project/infrastructure/config"
	"final-project/infrastructure/database"
	"final-project/repository/category_repo/category_pg"
	"final-project/repository/user_repo/user_pg"
	"final-project/service/auth_service"
	"final-project/service/category_service"
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

	categoryRepo := category_pg.NewCategoryRepo(db)
	categoryService := category_service.NewCategorySevice(categoryRepo)
	categoryHandler := category_handler.NewCategoryHandler(categoryService)

	authService := auth_service.NewAuthService(userRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.PUT("/update-account", authService.Authentication(), userHandler.Update)
		userRoute.DELETE("/delete-account", authService.Authentication(), userHandler.Delete)
	}

	userRoute = route.Group("/categories")
	{
		userRoute.POST("", authService.Authentication(), categoryHandler.Create)
	}
	route.Run(":" + config.AppConfig().Port)
}
