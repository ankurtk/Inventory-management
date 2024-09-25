package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"github.com/ankurtk/Inventory-management/Backend/models"
	"github.com/ankurtk/Inventory-management/Backend/database"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")


	pageInt, err := strconv.Atoi(page)
	if err!=nil || pageInt < 1 {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}

	offset := (pageInt - 1) * limitInt

	if err := database.DB.Offset(offset).Limit(limitInt).Find(&products).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"page":    pageInt,
		"limit":   limitInt,
		"products": products,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	for _, product := range products {
		// Check for duplicates by querying the database
		var existingProduct models.Product
		if err := database.DB.Where("name = ?", product.Name).First(&existingProduct).Error; err == nil {
			// If no error, it means a product with the same name exists
			http.Error(w, "Duplicate product name: "+product.Name, http.StatusConflict)
			return
		}

		// Create the product if it doesn't exist
		if err := database.DB.Create(&product).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(products)
}

func GetProductByID (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)  // getting id from URL
	idStr := vars["id"] // storing the in variable

	id, err := strconv.Atoi(idStr) // converting the id into int

	if(err!=nil || id<0){                                    // error handling
		http.Error(w, err.Error(),http.StatusBadRequest)
		return
	}

	// query in database for product with id

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {   // Gorm First method is used to query the product by id
        if errors.Is(err, gorm.ErrRecordNotFound) {
            response := map[string]interface{}{						// response creation
                "success": false,
                "message": "Product not found",
            }
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        } else {
            response := map[string]interface{}{
                "success": false,
                "message": "Internal server error",
            }
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(response)
        }
        return
    }
	response := map[string]interface{}{
        "success": true,
        "product": product,
    }


	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var existingProduct models.Product
	if err := database.DB.First(&existingProduct, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updatedProduct.Name != "" {
		existingProduct.Name = updatedProduct.Name
	}
	if updatedProduct.Category != "" {
		existingProduct.Category = updatedProduct.Category
	}
	if updatedProduct.Description != "" {
		existingProduct.Description = updatedProduct.Description
	}
	if updatedProduct.Quantity != nil {
		existingProduct.Quantity = updatedProduct.Quantity
	}

	if err := database.DB.Save(&existingProduct).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingProduct)
}

func DeleteProduct(w http.ResponseWriter,r*http.Request){
	var product models.Product
	vars:=mux.Vars(r)
	idStr:=vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := database.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err:=database.DB.Delete(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


// You can add more controllers for DeleteProduct
