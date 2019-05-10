package external_interfaces

import (
	"log"

	"github.com/jinzhu/gorm"
	"pumpkin/domain/model"
)

type database struct {
	db *gorm.DB
}

// create database
func NewDBOutput(dialect string, args ...interface{}) Output {
	db, err := gorm.Open(dialect, args)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &database{db}
}

func (d *database) Push(data interface{}) error {
	// initialize
	db := *d.db
	// noinspection ALL
	defer db.Close()

	pUser := data.(*model.User)
	db.Create(pUser)
	return nil
}
