package controller

import (
	"AppointmentsAPI/classes"
	"AppointmentsAPI/config"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.String(200, "THIS IS THE HOMEPAGE OF MY API")
}

func GetAppointments(c *gin.Context) {
	var appointments []classes.Appointment
	config.DB.Find(&appointments)
	c.JSON(200, &appointments)
}

func CreateAppointment(c *gin.Context) {
	var appointment classes.Appointment
	err := c.BindJSON(&appointment)
	if err != nil {
		return
	}
	config.DB.Create(&appointment)
	c.JSON(200, &appointment)
}

func DeleteAppointment(c *gin.Context) {
	var appointment classes.Appointment
	config.DB.Where("id = ?", c.Param("id")).Delete(&appointment)
	c.JSON(200, &appointment)
}

func UpdateAppointment(c *gin.Context) {
	var appointment classes.Appointment
	config.DB.Where("id = ?", c.Param("id")).First(&appointment)
	err := c.BindJSON(&appointment)
	if err != nil {
		return
	}
	config.DB.Save(&appointment)
	c.JSON(200, &appointment)
}
