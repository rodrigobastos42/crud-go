package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errorMessage := internalerrors.NewBadRequestError(fmt.Sprintf("%s is not a valid id", id))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := uc.service.DeleteUser(id); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusAccepted)
}
