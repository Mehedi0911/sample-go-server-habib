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
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	w.WriteHeader(201)
	utils.SendData(w, newProduct, http.StatusCreated)
}
