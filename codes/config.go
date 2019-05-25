package codes

// TODO: Refactor

import (
	"log"
	"os"
)

type config struct {
	key   string
	value string
}

type Configs []*config

func GetConfigs() *Configs {
	DBType := os.Getenv("DB_TYPE")
	DBName := os.Getenv("DB_NAME")
	log.Println(DBType, DBName)
	if DBType == "" {
		DBType = "sqlite3"
		DBName = "development.sqlite3"
	}
	log.Println(DBType, DBName)
	return &Configs{
		{key: "DBType", value: DBType},
		{key: "DBHost", value: os.Getenv("DB_HOST")},
		{key: "DBName", value: DBName},
		{key: "DBUser", value: os.Getenv("DB_USER")},
		{key: "DBPass", value: os.Getenv("DB_PASS")},
	}
}
func (cs *Configs) Get(key string) string {
	for _, conf := range *cs {
		if conf.key == key {
			return conf.value
		}
	}
	return ""
}
