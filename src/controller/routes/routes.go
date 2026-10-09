package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigobastos42/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.FindUserById)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
