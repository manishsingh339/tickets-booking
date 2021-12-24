package store

import (
	"context"
	"tickets/booking/internal/db"
)

const (
	DB         = "ticket_booking"
	Collection = "ticket"
)

func CreateIssue(ticket Ticket) error {
	client, err := db.Get()
	if err != nil {
		return err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection(Collection)

	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), ticket)

	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
