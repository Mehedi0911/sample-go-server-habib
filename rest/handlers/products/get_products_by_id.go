package products

import (
	"fmt"
	"net/http"
	"sample-server/utils"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")

	pId, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		return
	}

	if product == nil {
		utils.SendData(w, fmt.Sprintf("Product not found with id %d", pId), http.StatusNotFound)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
