package store

import (
	"context"
	"log"
	"tickets/booking/internal/db"
	"tickets/booking/internal/handler/v1/bookings/store/ticket"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get() ([]ticket.Ticket, error) {
	filter := bson.D{{}}
	tickets := []ticket.Ticket{}

	client := db.Get()

	collection := client.Database(ticket.DBConfig.DB).Collection(ticket.DBConfig.DB)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return tickets, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := ticket.Ticket{}
		err := cur.Decode(&t)
		if err != nil {
			return tickets, err
		}
		tickets = append(tickets, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(tickets) == 0 {
		return tickets, mongo.ErrNoDocuments
	}

	log.Println("Success")
	log.Println(tickets)

	//Return success without any error.
	return tickets, nil
}
