package route

import (
	"../controller"
	"github.com/gorilla/mux"
)

func GetAllRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()

	routes.HandleFunc("/users/register", controller.Register).Methods("POST")
	routes.HandleFunc("/users/login", controller.Login).Methods("POST")
	routes.HandleFunc("/users/{userId}", controller.GetUserData).Methods("GET")
	routes.HandleFunc("/users/{userId}/products", controller.GetUserUploadedProducts).Methods("GET")
	routes.HandleFunc("/users/{userId}/orders", controller.GetAllOrders).Methods("GET")

	routes.HandleFunc("/cities", controller.GetAllCities).Methods("GET")

	routes.HandleFunc("/categories", controller.GetAllCategories).Methods("GET")

	routes.HandleFunc("/products", controller.UploadProduct).Methods("POST")
	routes.HandleFunc("/products", controller.GetAllAvailableProducts).Methods("GET")
	routes.HandleFunc("/products/{productId}", controller.GetOneProductDetails).Methods("GET")
	routes.HandleFunc("/products/img/{imageName}", controller.GetProductImage).Methods("GET")

	routes.HandleFunc("/orders", controller.PlaceOrder).Methods("POST")
	routes.HandleFunc("/orders/retrieve", controller.ConfirmProductRetrieval).Methods("POST")
	routes.HandleFunc("/orders/return", controller.ConfirmProductReturn).Methods("POST")
	routes.HandleFunc("/orders/cancellation", controller.ConfirmProductCancellation).Methods("POST")

	return routes
}
