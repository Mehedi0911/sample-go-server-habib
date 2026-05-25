package handlers

import (
	"net/http"
	"sample-server/database"
	"sample-server/utils"
)

type loginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	var loginData loginDto

	if err := utils.HandleDecode(r, &loginData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	findUser, err := database.FindUserByEmail(loginData.Email, loginData.Password)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	utils.SendData(w, findUser, http.StatusOK)

}
