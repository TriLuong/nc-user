package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func AddStudent(student Student) (Student, error) {
	collection := Client.Database("nc-student").Collection("students")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, error := collection.InsertOne(ctx, student)
	if error != nil {
		return Student{}, error
	}
	newID, error := createId("studentID")
	if error != nil {
		log.Println(error)
		return Student{}, error
	}
	filter := bson.M{"_id": result.InsertedID}
	update := bson.M{"$set": bson.M{"id": newID}}
	// upsert := true
	// options := options.UpdateOptions{Upsert: &upsert}
	resultUpdate := collection.FindOneAndUpdate(ctx, filter, update)
	if resultUpdate.Err() != nil {
		return Student{}, resultUpdate.Err()
	}
	var studentResult Student
	resultUpdate.Decode(&studentResult)
	return studentResult, nil
}

func createId(name string) (int, error) {
	collection := Client.Database("nc-student").Collection("counters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// _, error := collection.InsertOne(ctx, bson.M{"_id": name})
	// if error != nil {
	// 	return 0, error
	// }
	var counter Counters
	filter := bson.M{"_id": name}
	update := bson.M{"$inc": bson.M{"id": 1}}
	upsert := true
	options := options.UpdateOptions{Upsert: &upsert}
	_, error := collection.UpdateOne(ctx, filter, update, &options)
	if error != nil {
		return 0, nil
	}
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return 0, result.Err()
	}
	error = result.Decode(&counter)
	if error != nil {
		log.Println(error)
		return 0, nil
	}
	log.Println(counter)

	return counter.ID, nil
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
