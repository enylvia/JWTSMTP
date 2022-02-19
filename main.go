package main

import (
	"github.com/gin-gonic/gin"
	"jwtsmtp/app"
	"jwtsmtp/handler"
	"jwtsmtp/repository"
	"jwtsmtp/service"
)

func main() {

	db := app.InitConnect()
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository)
	authService := service.NewJWTService()
	userHandler := handler.NewUserHandler(userService,authService)

	router := gin.Default()

	router.POST("/register", userHandler.RegisterUser)
	router.POST("/login", userHandler.Login)
	router.GET("/send", app.IsAuthorized(authService,userService),userHandler.SendMail)


	router.Run(":8080")
}

