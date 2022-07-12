package controller

import (
	"AppointmentsAPI/classes"
	"AppointmentsAPI/config"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HomePage(c *gin.Context) {
	c.String(200, "THIS IS THE HOMEPAGE OF MY API")
}

func GetAppointments(c *gin.Context) {
	var appointments []classes.Appointment
	config.DB.Find(&appointments)
	if len(appointments) > 0 {
		c.JSON(200, &appointments)
	} else {
		c.String(418, "Appointments list is empty")
	}
}

func GetOneAppointment(c *gin.Context) {
	var appointment classes.Appointment
	err1 := config.DB.Where("id = ?", c.Param("id")).First(&appointment).Error
	err2 := c.BindJSON(&appointment)
	if err2 != nil {
		return
	}
	if (classes.Appointment{}) != appointment {
		c.JSON(200, &appointment)
	} else if errors.Is(err1, gorm.ErrRecordNotFound) {
		c.String(404, "The appointment with id "+c.Param("id")+" does not exist.")
	}
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
