package route

import (
	. "../di"
	. "github.com/gorilla/mux"
)

func GetAllRoutes() *Router {
	routes := NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()

	userController := InjectUserController()
	routes.HandleFunc("/users/register", userController.RegisterHandler).Methods("POST")
	routes.HandleFunc("/users/login", userController.LoginHandler).Methods("POST")
	routes.HandleFunc("/users/{userId}", userController.GetUserDataHandler).Methods("GET")

	cityController := InjectCityController()
	routes.HandleFunc("/cities", cityController.GetAllCitiesHandler).Methods("GET")

	categoryController := InjectCategoryController()
	routes.HandleFunc("/categories", categoryController.FindAllCategoriesHandler).Methods("GET")

	productController := InjectProductController()
	routes.HandleFunc("/products", productController.UploadProductHandler).Methods("POST")
	routes.HandleFunc("/products", productController.GetAllAvailableProductsHandler).Methods("GET")
	routes.HandleFunc("/products/{productId}", productController.GetOneProductDetailsHandler).Methods("GET")
	routes.HandleFunc("/products/img/{imageName}", productController.GetProductImageHandler).Methods("GET")
	routes.HandleFunc("/users/{userId}/products", productController.GetUserUploadedProductsHandler).Methods("GET")

	orderController := InjectOrderController()
	routes.HandleFunc("/orders", orderController.PlaceOrderHandler).Methods("POST")
	routes.HandleFunc("/users/{userId}/orders", orderController.GetAllOrdersHandler).Methods("GET")
	routes.HandleFunc("/orders/retrieve", orderController.ConfirmProductRetrievalHandler).Methods("POST")
	routes.HandleFunc("/orders/return", orderController.ConfirmProductReturnHandler).Methods("POST")
	routes.HandleFunc("/orders/cancellation", orderController.ConfirmProductCancellationHandler).Methods("POST")

	return routes
}
