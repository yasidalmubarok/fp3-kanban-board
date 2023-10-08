package handler

import (
	"final-project/infrastructure/config"
	"final-project/infrastructure/database"
	"final-project/repository/user_repo/user_pg"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {

	config.LoadEnv()

	database.InitiliazeDatabase()
	db := database.GetDatabaseInstance()

	//Dependency Injection
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)

		userRoute.POST("/login", userHandler.Login)
	}

	route.Run(":" + config.AppConfig().Port)
}
