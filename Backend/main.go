package main

import (
    "log"
	"net/http"
	"github.com/ankurtk/Inventory-management/Backend/database"
	"github.com/ankurtk/Inventory-management/Backend/routes"
)

func main() {
	database.Connect()
	router := routes.RegisterRoutes()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
