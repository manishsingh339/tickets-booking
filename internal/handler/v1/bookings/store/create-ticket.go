package store

import (
	"context"
	"log"
	"tickets/booking/internal/db"
	"tickets/booking/internal/handler/v1/bookings/store/ticket"
)

func Save(_ticket *ticket.Ticket) error {
	client := db.Get()

	log.Println("Got DB Collection: ", _ticket)

	//Create a handle to the respective collection in the database.
	collection := client.Database(ticket.DBConfig.DB).Collection(ticket.DBConfig.DB)

	//Perform InsertOne operation & validate against the error.
	_, err := collection.InsertOne(context.TODO(), _ticket)

	if err != nil {
		return err
	}

	//Return success without any error.
	return nil
}
