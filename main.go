package main

import (
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
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	router.Run(":8181")
}
