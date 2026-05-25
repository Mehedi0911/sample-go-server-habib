package handlers

import (
	"net/http"
	"sample-server/database"
	"sample-server/utils"
)

func HandleCreateProducts(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	if err := utils.HandleDecode(r, &newProduct); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	prd := database.CreateProduct(newProduct)
	if prd == nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	utils.SendData(w, prd, http.StatusCreated)
}
