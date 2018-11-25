package helper

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)

type IDatabaseConnectionHelper interface {
	OpenDatabaseConnection() *gorm.DB
}

type DatabaseConnectionHelper struct {}

func (databaseConnectionHelper *DatabaseConnectionHelper) OpenDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(
		"postgres",
		"host = localhost " +
			"port = 5432 " +
			"user = postgres " +
			"dbname = seelih_dev " +
			"password = postgres " +
			"sslmode=disable")
	if err != nil {
		golog.Warn("Error connecting to DB : ", err)
	}
	return db
}