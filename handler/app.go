package handler

import (
	"final-project/handler/user_handler"
	"final-project/infrastructure/config"
	"final-project/infrastructure/database"
	"final-project/repository/user_repo/user_pg"
	"final-project/service/auth_service"
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

	authService := auth_service.NewAuthService(userRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.PUT("/update-account", authService.Authentication(), userHandler.Update)
	}

	route.Run(":" + config.AppConfig().Port)
}
