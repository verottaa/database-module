package database_module

func InitDatabase(uri string, dbName string) Code {
	return initialize(uri, dbName)
}
