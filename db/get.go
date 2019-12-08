package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Age       int                `json:"age" bson:"age"`
	Email     string             `json:"email" bson:"email"`
}

func GetStudents() (*[]Student, error) {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}
	return &students, nil
}

func GetStudentByID(id string) (interface{}, error) {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": newId}
	result := Student{}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
