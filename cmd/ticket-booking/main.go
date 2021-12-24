package main

import (
	"log"
	"net/http"
	"tickets/booking/internal/config"
	"tickets/booking/internal/db"
	"tickets/booking/internal/handler/v1/bookings"
)

func main() {
	log.Println("Connect to DB")
	dbConfig := config.GetDBConfig()
	db.Init(&db.Config{dbConfig.URL, false})
	http.HandleFunc("/ticket", bookings.Test)
	log.Println("Listing for requests at http://localhost:8000/ticket")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
