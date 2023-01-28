package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionTimeout = 10 * time.Second

func mongoDB(URI string) *mongo.Client {
	c, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatalf("Error while creating mongo client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	err = c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func newInstance(URI, dbName string) *mongo.Database {
	dbInstance := mongoDB(URI)
	return dbInstance.Database(dbName)
}
