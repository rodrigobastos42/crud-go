package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/configuration/validation"
	"github.com/rodrigobastos42/crud-go/src/controller/model/request"
	"github.com/rodrigobastos42/crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", "UpdateUser"))
	var userRequest *request.UserUpdateRequest

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		error := internalerrors.NewBadRequestError("id must be a hex value")
		c.JSON(error.Code, error)
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(id) == "" {
		logger.Error("Error on validate user info", err, zap.String("journey", "UpdateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain := model.NewUpdateUserDomain(
		userRequest.Name,
		userRequest.Age,
	)
	err := uc.service.UpdateUser(id, userDomain)
	if err != nil {
		c.JSON(err.Code, err)
	}
	logger.Info("User created successfully", zap.String("journey", "UpdateUser"))
	c.Status(http.StatusAccepted)
}
