package products

import (
	"net/http"
	"sample-server/repo"
	"sample-server/utils"
)

type RequestProductDTO struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

func (h *Handler) HandleCreateProducts(w http.ResponseWriter, r *http.Request) {

	var req RequestProductDTO

	if err := utils.HandleDecode(r, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	prd, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURL:      req.ImgURL,
	})
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	utils.SendData(w, prd, http.StatusCreated)
}
