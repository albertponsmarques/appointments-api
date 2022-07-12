package routes

import (
	"AppointmentsAPI/controller"
	"github.com/gin-gonic/gin"
)

func MyRoutes(r *gin.Engine) {
	r.GET("/", controller.HomePage)
	r.GET("/appointments", controller.GetAppointments)
	r.POST("/appointments", controller.CreateAppointment)
	r.DELETE("appointments/:id", controller.DeleteAppointment)
	r.PUT("appointments/:id", controller.UpdateAppointment)
}
