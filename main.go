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
	userHandler := handler.NewHandler(userService)

	router := gin.Default()
	v1 := router.Group("v1")
	v1.POST("/user", userHandler.Save)
	router.Run()
}
