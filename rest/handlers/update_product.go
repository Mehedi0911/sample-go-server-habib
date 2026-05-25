package handlers

import (
	"fmt"
	"net/http"
	"sample-server/database"
	"sample-server/utils"
	"strconv"
)

func HandleUpdateProducts(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")

	pId, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received update request for product ID: %d\n", pId)

	var newProduct database.Product

	if err := utils.HandleDecode(r, &newProduct); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	prd := database.UpdateProduct(pId, newProduct)
	if prd == nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	utils.SendData(w, prd, http.StatusCreated)
}
