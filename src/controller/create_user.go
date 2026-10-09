package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/configuration/validation"
	"github.com/rodrigobastos42/crud-go/src/controller/model/request"
	"github.com/rodrigobastos42/crud-go/src/model"
	"github.com/rodrigobastos42/crud-go/src/view"
	"go.uber.org/zap"
)

var UserDomainInterface model.UserDomainInterface

func (uc *userController) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))
	var userRequest *request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error on validate user info", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.CreateUser(userDomain); err != nil {
		c.JSON(err.Code, err)
	}
	logger.Info("User created successfully", zap.String("journey", "CreateUser"))
	c.JSON(http.StatusAccepted, view.ConvertDomainToResponse(userDomain))
}
