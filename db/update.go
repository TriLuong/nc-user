package db

import (
	"context"
	"fmt"
	"log"
	"time"
)

func AddStudent(student Student) {
	collection := Client.Database("nc-student").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	InsertResult, error := collection.InsertOne(ctx, student)
	if error != nil {
		log.Fatalln(error)
	}
	fmt.Println(InsertResult)
}
