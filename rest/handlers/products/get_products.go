package products

import (
	"net/http"
	"sample-server/utils"
)

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	list, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, list, http.StatusOK)
}
