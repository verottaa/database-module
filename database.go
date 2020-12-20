package database_module

import "go.mongodb.org/mongo-driver/bson/primitive"

func InitDatabase(uri string, dbName string) Code {
	return initialize(uri, dbName)
}

func CloseDatabase() Code {
	return closeDatabase()
}

func PushOne(collection string, object interface{}) (Code, interface{}) {
	return create(collection, object)
}

func FindOne(collection string, filter interface{}) (Code, interface{}) {
	return findOne(collection, filter)
}

func FindMany(collection string, filter interface{}) (Code, []interface{}) {
	return findMany(collection, filter)
}

func UpdateOne(collection string, filter interface{}, updatedObject interface{}) Code {
	return updateOne(collection, filter, updatedObject)
}

func DeleteOne(collection string, filter interface{}) Code {
	return deleteOne(collection, filter)
}

func DeleteMany(collection string, filter interface{}) (Code, int64) {
	return deleteMany(collection, filter)
}

func GenerateObjectID() primitive.ObjectID {
	return generateObjectID()
}
