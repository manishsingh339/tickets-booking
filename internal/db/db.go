package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

type Config struct {
	URL           string
	EnableDatadog bool
}

func NewDB(config *Config) (*mongo.Client, error) {
	opts := options.Client()
	if config.EnableDatadog {
		opts.Monitor = mongotrace.NewMonitor()
	}
	opts.ApplyURI(config.URL)
	c, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return c, nil
}

func Init(config *Config) error {
	d, err := NewDB(config)
	if err != nil {
		log.Println("Error while connecting to DB")
		return err
	}
	log.Println("Connected with DB")

	db = d
	return nil
}

func Close() error {
	return db.Disconnect(context.Background())
}

func Get() *mongo.Client {
	return db
}
