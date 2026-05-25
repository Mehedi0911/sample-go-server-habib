package handlers

import (
	"net/http"
	"sample-server/database"
	"sample-server/utils"
	"strconv"
)

func HandleDeleteProductByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")

	pId, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	database.DeleteProduct(pId)

	utils.SendData(w, "Deleted Successfully", http.StatusOK)
}
