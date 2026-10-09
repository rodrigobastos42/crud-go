package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigobastos42/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/users", controller.CreateUser)
	r.GET("/users/:id", controller.FindUserById)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
