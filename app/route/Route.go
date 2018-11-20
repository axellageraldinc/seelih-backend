package route

import (
	"../controller"
	"github.com/gorilla/mux"
)

func GetAllRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()

	routes.HandleFunc("/products", controller.GetAllAvailableProducts).Methods("GET")

	return routes
}