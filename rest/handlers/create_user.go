package handlers

import (
	"net/http"
	"sample-server/database"
	"sample-server/utils"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User

	if err := utils.HandleDecode(r, &newUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	insertedUser := newUser.Create()
	if insertedUser == nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	utils.SendData(w, insertedUser, http.StatusCreated)
}
