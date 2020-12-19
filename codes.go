package database_module

type Code int

const (
	SUCCESS        Code = 001
	ERROR          Code = 002
	DATABASE_EMPTY Code = 003
	CREATED        Code = 004
	FOUND          Code = 005
	NOT_FOUND      Code = 006
	FOUND_ANY      Code = 007
)
