package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userController) FindUserById(c *gin.Context) {
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errorMessage := internalerrors.NewBadRequestError(fmt.Sprintf("%s is not a valid id", id))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(*userDomain))
}
