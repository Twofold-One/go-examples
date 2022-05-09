package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongodbExample() {

	// create mongo.Client with Connect function
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:secret@localhost:27017/person"))

	// defer call to disconnect
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// ping the server
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	// TODO

	// // create the client
	// client, err := mongo.NewClient(`mongodb://root:secret@localhost:27017/person`)
	// if err != nil {
	// 	panic(err)
	// }

	// // connect to db
	// ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// // ping the server
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	panic(err)
	// }

}