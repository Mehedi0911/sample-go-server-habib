package utils

import (
	"encoding/json"
	"net/http"
	"sample-server/database"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func HandleDecode(r *http.Request, newProduct *database.Product) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&newProduct)
}
