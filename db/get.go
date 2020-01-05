package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

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

func GetStudentByID(id string) (Student, error) {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idInt, error := strconv.Atoi(id)
	if error != nil {
		return Student{}, error
	}
	filter := bson.M{"id": idInt}
	var result Student
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return Student{}, err
	}
	return result, nil
}

func DeleteStudentById(id string) (Student, error) {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idInt, error := strconv.Atoi(id)
	if error != nil {
		return Student{}, error
	}
	filter := bson.M{"id": idInt}
	var result Student
	err := collection.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return Student{}, err
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
