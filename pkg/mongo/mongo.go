package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoConnection(url, database, collection string) (conn *mongo.Collection, err error) {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(url)
	client, err1 := mongo.Connect(ctx, clientOptions)
	if err1 != nil {
		err = err1
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	conn = client.Database(database).Collection(collection)

	return
}
