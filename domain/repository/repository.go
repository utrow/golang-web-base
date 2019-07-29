package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/utrow/golang-web-base/infrastructure/database/repository"
)

type Repository struct {
	DB databaseRepository.DatabaseRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		DB: databaseRepository.NewDatabaseRepository(db),
	}
}
