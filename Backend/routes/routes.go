package routes

import (
	"github.com/gorilla/mux"
	"github.com/ankurtk/Inventory-management/Backend/controllers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.GetProductByID).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	return router
}
