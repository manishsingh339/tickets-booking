package main

import (
	"log"
	"net/http"
	"tickets/booking/internal/config"
	"tickets/booking/internal/db"
	bookings "tickets/booking/internal/handler/v1/bookings/manager"
)

func main() {
	initConfigs()

	http.HandleFunc("/ping", bookings.Ping)
	http.HandleFunc("/ticket", bookings.AddTicket)
	http.HandleFunc("/tickets", bookings.Tickets)
	log.Println("Listing for requests at http://localhost:8000/ticket")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initConfigs() {
	log.Println("Connect to DB")
	dbConfig := config.GetDBConfig()
	db.Init(&db.Config{dbConfig.URL, false})
}
