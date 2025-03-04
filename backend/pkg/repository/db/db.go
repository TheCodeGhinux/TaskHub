package db

import (
	"gorm.io/gorm"
)

type Database struct {
	Postgres *gorm.DB
}

var (
	DB *Database = &Database{}
)

func Connect() *Database {
	return DB
}
