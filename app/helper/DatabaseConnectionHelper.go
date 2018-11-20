package helper

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)

func OpenDatabaseConnection() *gorm.DB {
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
	} else {
		golog.Info("Success connecting to DB!")
	}
	return db
}