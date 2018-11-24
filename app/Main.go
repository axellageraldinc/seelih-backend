package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"../app/controller"
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
	// data for cities and categories are in the dummy_default_data.sql file

	// USER'S DEFAULT DATA
	user1 := model.User{
		Phone: "08123123123",
		Fulladdress: "Yogyakarta",
		Fullname: "Axellageraldinc Adryamarthanino",
		CityCodeId: 123,
		Password: controller.HashAndSalt(controller.ConvertPlainPasswordToByte("axell123")),
		Email: "axell@gmail.com",
	}
	user2 := model.User{
		Phone: "08456456456",
		Fulladdress: "Lamongan",
		Fullname: "Moh Azzum Jordhan Wiratama",
		CityCodeId: 111,
		Password: controller.HashAndSalt(controller.ConvertPlainPasswordToByte("azzum123")),
		Email: "azzum@gmail.com",
	}
	user3 := model.User{
		Phone: "08678678678",
		Fulladdress: "Boyolali",
		Fullname: "Almantera Tiantana Al Faruqi",
		CityCodeId: 222,
		Password: controller.HashAndSalt(controller.ConvertPlainPasswordToByte("alman123")),
		Email: "alman@gmail.com",
	}
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)

	// PRODUCT'S DEFAULT DATA
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
