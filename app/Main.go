package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"

	//"../app/controller"
	"../app/helper"
	"../app/model"
	"../app/route"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

func main() {
	initiateMigration()
	initiateRoutes()
}

func initiateMigration() {
	db := helper.OpenDatabaseConnection()
	defer db.Close()

	// db.DropTable(&model.Category{}, &model.City{}, &model.User{}, &model.Product{}, &model.Order{})   // Uncomment this code if there's a column deletion in DB
	db.AutoMigrate(&model.Category{}, &model.City{}, &model.User{}, &model.Product{}, &model.Order{}) // WILL create table, add missing columns, WON'T change column type/delete column
	// insertDefaultData(db)
}

func initiateRoutes() {
	routes := route.GetAllRoutes()
	http.Handle("/", routes)
	handler := cors.Default().Handler(routes)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func insertDefaultData(db *gorm.DB) {
	product1 := model.Product{
		TenantID:            1,
		CategoryID:          4,
		ImageName:           "camera.png",
		Quantity:            1,
		PricePerItemPerDay:  20000,
		ProductStatus:       model.OPENED,
		MinimumBorrowedTime: 1,
		MaximumBorrowedTime: 3,
		Description:         "mirrorless camera by SONY",
		Sku:                 "cam_mirrorless_123",
		Name:                "Sony Mirrorless Camera",
	}
	product2 := model.Product{
		TenantID:            2,
		CategoryID:          6,
		ImageName:           "drone.jpg",
		Quantity:            1,
		PricePerItemPerDay:  200000,
		ProductStatus:       model.OPENED,
		MinimumBorrowedTime: 1,
		MaximumBorrowedTime: 1,
		Description:         "drone with camera by DJI",
		Sku:                 "elec_drone_dji_phantom_123",
		Name:                "Dji Phantom",
	}
	db.Create(&product1)
	db.Create(&product2)
}
