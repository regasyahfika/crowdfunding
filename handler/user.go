package handler

import (
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h userHandler) Save(c *gin.Context) {
	var userRequest user.RegisterUserInput

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		log.Fatal("error")
	}
}
