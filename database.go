package database_module

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

func FindMany(collection string, filter interface{}) (Code, interface{}) {
	return findMany(collection, filter)
}

func Update(filter interface{}, object interface{}) {

}

func DeleteOne(filter interface{}) {

}

func DeleteAll(filter interface{}) {

}
