package store

import (
	"context"
	"log"
	"tickets/booking/internal/db"
	"tickets/booking/internal/handler/v1/bookings/store/ticket"
)

const (
	DB         = "ticket_booking"
	Collection = "ticket"
)

func Save(ticket *ticket.Ticket) error {
	client := db.Get()

	log.Println("Got DB Collection: ", ticket)

	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Collection)

	//Perform InsertOne operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), ticket)

	if err != nil {
		return err
	}

	//Return success without any error.
	return nil
}
