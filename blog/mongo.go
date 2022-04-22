package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Number struct {
	Name  string  `bson:"name"` //used for parsting from mongo
	Value float64 `bson:"value"`
}

func MongoInsert() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID

	fmt.Println(id)
}

func MongoInsertObject() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	num := Number{
		Name:  "epsilon",
		Value: 2.789,
	}

	res, err := collection.InsertOne(ctx, num)
	id := res.InsertedID

	fmt.Println(id)
}

func MongoGetList() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	var num Number

	res, err := collection.Find(ctx, bson.D{{}})
	for res.Next(ctx) { //iterator
		res.Decode(&num)
		elems, _ := res.Current.Elements()
		fmt.Println(elems[0])
		fmt.Println(num)
	}
}

func MongoGetOne() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	var num Number

	res := collection.FindOne(ctx, bson.D{{"name", "pi"}})
	res.Decode(&num)
	fmt.Println(num)
}

func main() {
	// MongoInsert()
	// MongoInsertObject()
	// MongoGetList()
	MongoGetOne()
}
