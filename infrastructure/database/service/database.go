package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"os"
)

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	DBNAME := os.Getenv("DB_DATABASE")
	OPTION := "charset=utf8mb4,utf8&parseTime=True"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

func CreateConnection() (*gorm.DB, error) {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		return nil, errors.New("DBコネクションの生成に失敗しました")
	}

	if os.Getenv("APP_ENV") == "local" {
		db.Debug()
	}

	return db, nil
}
