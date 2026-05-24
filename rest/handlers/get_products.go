package handlers

import (
	"net/http"
	"sample-server/database"
	"sample-server/utils"
)

func HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ProductList, http.StatusOK)
}
