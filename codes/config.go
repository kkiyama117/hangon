package codes

// TODO: Refactor

import (
	"os"
)

type config struct {
	Key   string
	Value string
}

type Configs []*config

func GetConfigs() *Configs {
	DBType := os.Getenv("DB_TYPE")
	DBName := os.Getenv("DB_NAME")
	if DBType == "" {
		DBType = "sqlite3"
		DBName = "development.sqlite3"
	}
	return &Configs{
		{Key: "DBType", Value: DBType},
		{Key: "DBHost", Value: os.Getenv("DB_HOST")},
		{Key: "DBName", Value: DBName},
		{Key: "DBUser", Value: os.Getenv("DB_USER")},
		{Key: "DBPass", Value: os.Getenv("DB_PASS")},
	}
}
func (cs *Configs) GetValue(key string) string {
	for _, conf := range *cs {
		if conf.Key == key {
			return conf.Value
		}
	}
	return ""
}

func (cs *Configs) Get(key string) (string, string) {
	for _, conf := range *cs {
		if conf.Key == key {
			return conf.Key, conf.Value
		}
	}
	return "", ""
}
