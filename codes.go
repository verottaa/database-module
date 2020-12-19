package database_module

type Code string

const (
	SUCCESS        Code = "001"
	ERROR          Code = "002"
	DATABASE_EMPTY Code = "003"
	CREATED        Code = "004"
)
