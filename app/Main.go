package main

import (
	"../app/model"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

func main() {
	connectToDatabase()
}

func connectToDatabase()  {
	db, err := gorm.Open(
		"postgres",
		"host = localhost " +
			"port = 5432 " +
			"user = postgres " +
			"dbname = seelih_dev " +
			"password = postgres " +
			"sslmode=disable")
	//db.DropTable(&model.Category{}, &model.City{}, &model.User{}, &model.Product{}, &model.Order{}) // Uncomment this code if there's a column deletion in DB
	db.AutoMigrate(&model.Category{}, &model.City{}, &model.User{}, &model.Product{}, &model.Order{}) // WILL create table, add missing columns, WON'T change column type/delete column
	if err != nil {
		golog.Warn("Error connecting to DB : ", err)
	} else {
		golog.Info("Success connecting to DB!")
	}
	defer db.Close()
}
