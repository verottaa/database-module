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
	_client             *mongo.Client
)

func initialize(uri string, dbName string) Code {
	setURI(uri)
	setDatabaseName(dbName)

	createContext()
	createClientDatabaseOptions()

	err := createClient()
	if err != nil {
		return ERROR
	}

	return ping()
}

func closeDatabase() Code {
	err := _client.Disconnect(_context)
	if err != nil {
		return ERROR
	}
	return SUCCESS
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

func createClient() error {
	client, err := mongo.Connect(_context, _mongoClientOptions)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_client = client
	return nil
}

func ping() Code {
	db := _client.Database(_databaseName)
	if db.Name() == _databaseName {
		return SUCCESS
	}
	return DATABASE_EMPTY
}
