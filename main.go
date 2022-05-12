package main

import (
	"crowdfunding/auth"
	"crowdfunding/config"
	"crowdfunding/handler"
	"crowdfunding/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := config.Connect(&gorm.DB{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run(":8181")
}
