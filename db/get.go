package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	// MongoID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID        int    `json:"id,omitempty" bson:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	Age       int    `json:"age,omitempty" bson:"age"`
	Email     string `json:"email,omitempty" bson:"email"`
}

type StudentSearchRequest struct {
	// ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Counters struct {
	NameID string `json:"_id,omitempty" bson:"_id,omitempty"`
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
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

func SerchStudentSimple(req StudentSearchRequest) (*[]Student, error) {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var filter bson.M
	bs, err := json.Marshal(req)

	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}
	err = json.Unmarshal(bs, &filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}
	fmt.Println(filter)
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
