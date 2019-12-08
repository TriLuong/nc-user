package db

import (
	"context"
	"time"
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
