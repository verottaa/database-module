package database_module

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	_uri                string
	_databaseName       string
	_context            context.Context
	_mongoClientOptions *options.ClientOptions
)

func initialize(uri string, dbName string) Code {
	setURI(uri)
	setDatabaseName(dbName)

	createContext()
	createClientDatabaseOptions()

	return ping()
}

func setURI(uri string) {
	_uri = uri
}

func setDatabaseName(name string) {
	_databaseName = name
}

func createContext() {
	_context = context.Background()
}

func createClientDatabaseOptions() {
	_mongoClientOptions = options.Client().ApplyURI(_uri)
}

func ping() Code {
	client, err := mongo.Connect(_context, _mongoClientOptions)
	if err != nil {
		fmt.Println(err)
		return ERROR
	}

	db := client.Database(_databaseName)
	fmt.Println(db.Name())
	if db.Name() == _databaseName {
		return SUCCESS
	}
	return DATABASE_EMPTY
}
