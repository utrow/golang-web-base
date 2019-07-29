package databaseRepository

import (
	"github.com/jinzhu/gorm"
)

type DatabaseRepository interface {
}

type db struct {
	db     DatabaseRepository
	dbConn *gorm.DB
}

func NewDatabaseRepository(d *gorm.DB) DatabaseRepository {
	return &db{
		dbConn: d,
	}
}
