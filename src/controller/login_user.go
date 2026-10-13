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

func (uc *userController) Login(c *gin.Context) {
	logger.Info("Init login controller", zap.String("journey", "login"))
	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error trying to login", err, zap.String("journey", "login"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := model.NewUserLoginDomain(userLogin.Email, userLogin.Password)

	result, token, err := uc.service.Login(userDomain)

	if err != nil {
		logger.Error("Error trying to call login Service", err, zap.String("journey", "login"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Login controller executed successfuly", zap.String("userId", result.GetID()), zap.String("journey", "login"))
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))
}
