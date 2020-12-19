package database_module

import (
	"context"
	"fmt"
	"time"
)

func create(collectionName string, object interface{}) (Code, interface{}) {
	collection := _client.Database(_databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(_context, 1*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, object)
	if err != nil {
		fmt.Println(err.Error())
		return ERROR, nil
	}

	return SUCCESS, result.InsertedID
}
