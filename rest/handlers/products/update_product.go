package products

import (
	"fmt"
	"net/http"
	"sample-server/repo"
	"sample-server/utils"
	"strconv"
)

type UpdateProductRequestDTO struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

func (h *Handler) HandleUpdateProducts(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")

	pId, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received update request for product ID: %d\n", pId)

	var req UpdateProductRequestDTO

	if err := utils.HandleDecode(r, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	updateErr := h.productRepo.Update(&repo.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURL:      req.ImgURL,
	})
	if updateErr != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	utils.SendData(w, fmt.Sprintf("Product updated successfully"), http.StatusCreated)
}
