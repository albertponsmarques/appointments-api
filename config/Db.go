package config

import (
	"AppointmentsAPI/classes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://awotxvic:uxJrwNzW7MYyb9BaJUThmdVAIP0fJ4jr@surus.db.elephantsql.com/awotxvic"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&classes.Appointment{})
	if err != nil {
		return
	}

	DB = db
}
