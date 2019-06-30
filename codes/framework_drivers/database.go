package framework_drivers

import (
	"github.com/jinzhu/gorm"
	"pumpkin/codes/interface_adapters/db/_store_user/driver_ports"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	"pumpkin/codes/domain/model"
)

type database struct {
	db *gorm.DB
}

// create database
func NewDBOutput(dbFunc func() *gorm.DB) driver_ports.DBOutput {
	return &database{db: dbFunc()}
}

func (d *database) StoreUser(pUser *model.User) error {
	// initialize
	db := *d.db
	// noinspection ALL
	defer db.Close()

	db.Create(pUser)
	db.Where("user_name=?", pUser.UserName).First(pUser)
	return nil
}
