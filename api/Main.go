package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

// Main function
func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all appointments
	router.HandleFunc("/appointments/", GetAppointments).Methods("GET")

	// Create an appointment
	router.HandleFunc("/appointments/", CreateAppointment).Methods("POST")

	// Delete a specific appointment by the appointmentID
	router.HandleFunc("/appointments/{appointmentid}", DeleteAppointment).Methods("DELETE")

	// Delete all appointments
	router.HandleFunc("/appointments/", DeleteAppointments).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting appointments...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM appointments")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var appointments []Appointment

	// Foreach movie
	for rows.Next() {
		var id int
		var name string
		var date time.Time
		var hour time.Time

		err = rows.Scan(&id, &name, &date, &hour)

		// check errors
		checkErr(err)

		appointments = append(appointments, Appointment{Id: id, Nom: name, Dia: date, Hora: hour})
	}

	var response = JsonResponse{Type: "success", Data: appointments}

	json.NewEncoder(w).Encode(response)
}

//si
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := r.FormValue("appointmentid")
	appointmentName := r.FormValue("appointmentname")
	appointmentDate := r.FormValue("appointmentDate")
	appointmentHour := r.FormValue("appointmentHour")

	var response = JsonResponse{}

	if appointmentId == "" || appointmentName == "" || appointmentDate == "" || appointmentHour == "" {
		response = JsonResponse{Type: "error", Message: "You are missing parameters."}
	} else {
		db := setupDB()

		printMessage("Inserting appointment into DB")

		fmt.Println("Inserting new appointment with ID: " + appointmentId + " and name: " + appointmentName +
			" with date: " + appointmentDate + " at: " + appointmentHour)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO appointments(id, nom, dia, hora) VALUES($1, $2, $3, $4) returning id;", appointmentId, appointmentName, appointmentDate, appointmentHour).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The appointment has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	appointmentID := params["appointmentid"]

	var response = JsonResponse{}

	if appointmentID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing the ID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting appointment from DB")

		_, err := db.Exec("DELETE FROM appointments where id = $1", appointmentID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The appointment has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteAppointments(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all appointments...")

	_, err := db.Exec("DELETE FROM appointments")

	// check errors
	checkErr(err)

	printMessage("All appointments have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All appointments have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// ! other
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//DataBase
const (
	DB_USER     = "awotxvic"
	DB_PASSWORD = "uxJrwNzW7MYyb9BaJUThmdVAIP0fJ4jr"
	DB_NAME     = "appointments"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

// Classes

type Appointment struct {
	Id   int       `json:"appointmentid"`
	Nom  string    `json:"appointmentname"`
	Dia  time.Time `json:"appointmentdate"`
	Hora time.Time `json:"appointmenthour"`
}

type JsonResponse struct {
	Type    string        `json:"type"`
	Data    []Appointment `json:"data"`
	Message string        `json:"message"`
}
