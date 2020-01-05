package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Triluong/nc-student/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

var Client *mongo.Client

func Test() {
	InsertNumber()
	fmt.Println("Connected to DB")
}

func Init() {
	connect()
	// InsertNumber()
}

func InsertNumber() {
	collection := Client.Database("testing").Collection("numbers")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	id := res.InsertedID

	fmt.Println(id)
}

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.URI))

	if err != nil {
		log.Println("Error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Error")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("Error")
	}

	Client = client
	// fmt.Println("Connected to MongoDB!")
}
