package main

import (
	"AppointmentsAPI/config"
	"AppointmentsAPI/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.New()

	config.Connect()

	routes.MyRoutes(router)

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
