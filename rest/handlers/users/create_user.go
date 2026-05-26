package users

import (
	"net/http"
	"sample-server/repo"
	"sample-server/utils"
)

type CreateUserRequestDTO struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	var req CreateUserRequestDTO

	if err := utils.HandleDecode(r, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	insertedUser, err := h.userRepo.Create(repo.User{

		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	utils.SendData(w, insertedUser, http.StatusCreated)
}
