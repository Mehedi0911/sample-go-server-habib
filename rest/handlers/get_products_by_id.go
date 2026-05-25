package handlers

import (
	"fmt"
	"net/http"
	"sample-server/database"
	"sample-server/utils"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")

	pId, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.GetProductByID(pId)

	if product == nil {
		utils.SendData(w, fmt.Sprintf("Product not found with id %d", pId), http.StatusNotFound)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
