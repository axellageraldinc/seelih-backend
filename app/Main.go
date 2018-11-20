package main

import (
	"../app/helper"
	"../app/model"
	"../app/route"
	"log"
	"net/http"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

func main() {
	initiateMigration()
	initiateRoutes()
}

func initiateMigration()  {
	db := helper.OpenDatabaseConnection()
	defer db.Close()
	
	//db.DropTable(&model.Category{}, &model.City{}, &model.User{}, &model.Product{}, &model.Order{}) // Uncomment this code if there's a column deletion in DB
	db.AutoMigrate(&model.Category{}, &model.City{}, &model.User{}, &model.Product{}, &model.Order{}) // WILL create table, add missing columns, WON'T change column type/delete column
}

func initiateRoutes()  {
	routes := route.GetAllRoutes()
	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}