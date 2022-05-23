package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongodbExample() {

	// create mongo.Client with Connect function
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

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

	collection := client.Database("db").Collection("person")

	// insert document to collection
	// insertDocument(collection)

	// read document from collection
	readDocument(collection)

	// logical query
	logicalDocumentQuery(collection)

	// update document
	updateDocument(collection)
}

// insert document
func insertDocument(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "Jhon"}, {"lastname", "Doe"}})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)
}

// read a document
type Person struct {
	Name     string
	Lastname string
}

func readDocument(collection *mongo.Collection) {
	filter := bson.D{{"name", "Peter"}}
	result := Person{}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", result)
}

// query document with logical selector
func logicalDocumentQuery(collection *mongo.Collection) {
	q1 := bson.M{"name": bson.M{"$eq": "Jhon"}}
	q2 := bson.M{"lastname": bson.M{"$eq": "Doe"}}
	clauses := bson.A{q1, q2}
	filter := bson.M{"$and": clauses}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
}

// update a document
func updateDocument(collection *mongo.Collection) {
	filter := bson.M{"name": "Jhon"}
	update := bson.M{"$set": bson.M{"name": "John"}}

	res := collection.FindOneAndUpdate(context.Background(), filter, update)
	resDecoded := Person{}
	err := res.Decode(&resDecoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resDecoded)
}
