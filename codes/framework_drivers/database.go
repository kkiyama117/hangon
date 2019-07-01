package framework_drivers

import (
	"github.com/jinzhu/gorm"
	"pumpkin/codes/interface_adapters/db/user/driver_ports"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	"pumpkin/codes/domain/model"
)

type database struct {
	db *gorm.DB
}

// create database
func NewDBOutput(dbFunc func() *gorm.DB) driver_ports.UserDBOutput {
	return &database{db: dbFunc()}
}

func (d *database) GetUser(pUser *model.User) error {
	// initialize
	db := *d.db
	// noinspection ALL
	defer db.Close()

	db.Where(pUser).First(pUser)
	return nil
}


func (d *database) GetUsers(users model.Users) error {
	// initialize
	db := *d.db
	// noinspection ALL
	defer db.Close()

	db.Find(&users)
	return nil
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
