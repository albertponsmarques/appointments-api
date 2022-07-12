package classes

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	Id   int    `json:"ID" gorm:"primary_key"`
	Name string `json:"name"`
	Tlf  string `json:"tlf"`
	Day  string `json:"day"`
	Hour string `json:"hour"`
}
