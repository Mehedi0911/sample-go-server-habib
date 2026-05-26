package products

import (
	"net/http"
	"sample-server/utils"
	"strconv"
)

func (h *Handler) HandleDeleteProductByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")

	pId, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, "Deleted Successfully", http.StatusOK)
}
