package handlers

import (
	"net/http"
	"sample-server/config"
	"sample-server/database"
	"sample-server/utils"
	"strconv"
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

	user, err := database.FindUserByEmail(loginData.Email, loginData.Password)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	accessToken, err := utils.CreateJWT(config.GetConfig().JwtSecret, utils.Payload{
		Sub:         strconv.Itoa(user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusOK)

}
