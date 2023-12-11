package config

import (
	"os"
)

type DB struct {
	ConnStr string
}

var db = &DB{}

func DBCfg() *DB {
	return db
}

func LoadDB() {
	db.ConnStr = os.Getenv("DB_CONN_STR")
}
