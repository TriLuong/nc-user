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
