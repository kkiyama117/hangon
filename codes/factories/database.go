package factories

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"pumpkin/codes"
)

type dbOption struct {
	dialect string
	args    string
}

type DBOption func(option *dbOption)

// TODO: Refactor
func WithDBOption(configs *codes.Configs) DBOption {
	dialect := configs.GetValue("DBType")
	dbName := configs.GetValue("DBName")
	switch dialect {
	case "sqlite3":
		return func(option *dbOption) {
			option.dialect = dialect
			option.args = dbName
		}
	case "postgres":
		return func(option *dbOption) {
			option.dialect = dialect
			dbHost := configs.GetValue("DBHost")
			dbUser := configs.GetValue("DBUser")
			dbPass := configs.GetValue("DBPass")
			option.args = "host=" + dbHost + " user=" + dbUser + " db_name=" + dbName + " password=" + dbPass
		}
	default:
		log.Fatal("option is nil")
		return nil
	}
}

func InjectDBFunc(optionFunc ...DBOption) func() *gorm.DB {
	return func() *gorm.DB {
		opt := dbOption{}
		for _, o := range optionFunc {
			o(&opt)
		}
		db, err := gorm.Open(opt.dialect, opt.args)
		if err != nil {
			log.Print(err)
			return nil
		}
		return db
	}
}
