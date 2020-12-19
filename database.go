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

func FindMany(collection string, filter interface{}) (Code, []interface{}) {
	return findMany(collection, filter)
}

func Update(collection string, filter interface{}, object interface{}) (Code, interface{}) {
	return 0, nil
}

func DeleteOne(collection string, filter interface{}) Code {
	return deleteOne(collection, filter)
}

func DeleteAll(collection string, filter interface{}) (Code, int64) {
	return deleteMany(collection, filter)
}
