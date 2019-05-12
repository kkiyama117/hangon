package framework_drivers

import (
	"log"

	"github.com/jinzhu/gorm"

	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
	"pumpkin/domain/model"
	"pumpkin/interface_adapters/_store_user/driver_ports"
)

type database struct {
	db *gorm.DB
}

type dbOption struct {
	dialect string
	args    string
}
type DBOption func(option *dbOption)

func WithDBOption(dialect string, args string) DBOption {
	return func(option *dbOption) {
		option.dialect = dialect
		option.args = args
	}
}

// create database
func NewDBOutput(options ...DBOption) driver_ports.DBOutput {
	opt := dbOption{}
	for _, o := range options {
		o(&opt)
	}
	db, err := gorm.Open(opt.dialect, opt.args)
	if err != nil {
		log.Print(err)
		return nil
	}
	return &database{db}
}

func (d *database) StoreUser(pUser *model.User) error {
	// initialize
	db := *d.db
	// noinspection ALL
	defer db.Close()

	db.Create(pUser)
	db.First(pUser, "user_name", pUser.UserName)
	return nil
}
