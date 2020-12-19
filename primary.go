package database_module

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func getCollection(collectionName string) *mongo.Collection {
	return _client.Database(_databaseName).Collection(collectionName)
}

func create(collectionName string, object interface{}) (Code, interface{}) {
	collection := getCollection(collectionName)

	ctx, cancel := context.WithTimeout(_context, 1*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, object)
	if err != nil {
		fmt.Println(err.Error())
		return ERROR, nil
	}

	return CREATED, result.InsertedID
}

func findOne(collectionName string, filter interface{}) (Code, interface{}) {
	collection := getCollection(collectionName)

	ctx, cancel := context.WithTimeout(_context, 1*time.Second)
	defer cancel()

	var result interface{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return ERROR, nil
	}
	return FOUND, result
}

func findMany(collectionName string, filter interface{}) (Code, []interface{}) {
	collection := getCollection(collectionName)

	ctx, cancel := context.WithTimeout(_context, 1*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}()

	var resultArray []interface{}
	for cursor.Next(ctx) {
		var result interface{}
		err = cursor.Decode(&result)
		if err != nil {
			fmt.Println(err.Error())
		}
		resultArray = append(resultArray, result)
	}

	if len(resultArray) == 0 {
		return NOT_FOUND, resultArray
	}
	return FOUND_ANY, resultArray
}
