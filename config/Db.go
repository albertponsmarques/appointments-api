package config

import (
	"AppointmentsAPI/classes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://ukzhonijtpkwdc:b289a4d18e72984552bd10f251ed0e3a869a5bb9c9c3dcd4bc641f2bfd5e7428@ec2-54-228-32-29.eu-west-1.compute.amazonaws.com:5432/d2qmn2h3i98o9t"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&classes.Appointment{})
	if err != nil {
		return
	}

	DB = db
}
