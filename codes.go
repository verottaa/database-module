package database_module

type Code int

const (
	SUCCESS Code = iota
	ERROR
	DATABASE_EMPTY
	CREATED
	FOUND
	NOT_FOUND
	FOUND_ANY
	DELETED
)
