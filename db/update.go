package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func AddStudent(student interface{}) error {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, error := collection.InsertOne(ctx, student)
	if error != nil {
		return error
	}
	return nil
}

func UpdateStudent(id string, student Student) error {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newId, error := primitive.ObjectIDFromHex(id)
	if error != nil {
		return error
	}

	filter := bson.M{"_id": newId}
	update := bson.M{"$set": student}
	_, error = collection.UpdateOne(ctx, filter, update)
	if error != nil {
		return error
	}
	return nil
}
