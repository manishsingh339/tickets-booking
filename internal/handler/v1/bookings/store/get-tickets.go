package store

import (
	"context"
	"log"
	"tickets/booking/internal/db"
	"tickets/booking/internal/handler/v1/bookings/store/ticket"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketsRes struct {
	Tickets           []ticket.Ticket `json:"tickets" bson:"tickets"`
	CurrentSystemTime time.Duration   `json:"currentSystemTime" bson:"currentSystemTime"`
}

func Get() (TicketsRes, error) {
	filter := bson.D{}
	tickets := []ticket.Ticket{}
	ticketsRes := TicketsRes{nil, 0}

	client := db.Get()

	collection := client.Database(ticket.DBConfig.DB).Collection(ticket.DBConfig.Collection)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		log.Println("Error while find")
		return ticketsRes, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := ticket.Ticket{}
		err := cur.Decode(&t)
		if err != nil {
			log.Println("Error while Decode")
			return ticketsRes, err
		}
		log.Println("cur ", t)
		tickets = append(tickets, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(tickets) == 0 {
		log.Println("Error while len")
		return ticketsRes, mongo.ErrNoDocuments
	}

	ticketsRes = TicketsRes{Tickets: tickets}

	log.Println("Success")
	log.Println(tickets)

	//Return success without any error.
	return ticketsRes, nil
}
