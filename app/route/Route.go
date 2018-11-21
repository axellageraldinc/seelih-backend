package route

import (
	"../controller"
	"github.com/gorilla/mux"
)

func GetAllRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()

	routes.HandleFunc("/products", controller.GetAllAvailableProducts).Methods("GET")
	routes.HandleFunc("/products/{productId}", controller.GetOneProductDetails).Methods("GET")

	routes.HandleFunc("/orders", controller.PlaceOrder).Methods("POST")
	routes.HandleFunc("/orders/reception", controller.ConfirmProductRetrieval).Methods("POST")

	return routes
}