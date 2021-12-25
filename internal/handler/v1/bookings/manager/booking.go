package bookings

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	CreateTicket "tickets/booking/internal/handler/v1/bookings/store"
	GetTickets "tickets/booking/internal/handler/v1/bookings/store"
	"tickets/booking/internal/handler/v1/bookings/store/ticket"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	io.WriteString(w, "Pong!\n")
}

func AddTicket(w http.ResponseWriter, req *http.Request) {
	_ticket := ticket.Ticket{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		EventTime:   time.Now(),
		EventTitle:  "Event Title",
		Description: "Some description",
	}

	CreateTicket.Save(&_ticket)

	io.WriteString(w, "Ticket Created!\n")
}

func Tickets(w http.ResponseWriter, req *http.Request) {
	data, error := GetTickets.Get()
	if error != nil {
		panic("Error: Language cannot be nil")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
