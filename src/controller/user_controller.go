package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigobastos42/crud-go/src/model/service"
)

func NewUserController(service service.UserDomainService) UserControllerInterface {
	return &userController{service}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userController struct {
	service service.UserDomainService
}
