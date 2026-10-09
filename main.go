package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/controller"
	"github.com/rodrigobastos42/crud-go/src/controller/routes"
	"github.com/rodrigobastos42/crud-go/src/model/service"
)

func main() {
	logger.Info("application starting...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserController(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TEST"))
}
