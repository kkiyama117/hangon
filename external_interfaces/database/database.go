package database

type Database interface {
}

type database struct {
}

func NewDatabase() Database {
	return &database{}
}
