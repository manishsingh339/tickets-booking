package main

import (
	"log"
	"net/http"
	"tickets/booking/internal/handler/v1/bookings"
)

func main() {
	http.HandleFunc("/ticket", bookings.Test)
	log.Println("Listing for requests at http://localhost:8000/ticket")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
