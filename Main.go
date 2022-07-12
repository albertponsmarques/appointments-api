package main

import (
	"AppointmentsAPI/config"
	"AppointmentsAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.Connect()

	routes.MyRoutes(router)

	//err := router.Run(":" + os.Getenv("PORT"))
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
