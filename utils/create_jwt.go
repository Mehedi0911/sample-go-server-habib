package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg  string `json:"alg"`
	Type string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, payload Payload) (string, error) {
	header := Header{
		Alg:  "HS256",
		Type: "JWT",
	}

	headerJSON, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	payloadJSON, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	byteArrSecret := []byte(secret)

	message := base64URLEncode(headerJSON) + "." + base64URLEncode(payloadJSON)

	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)

	jwt := message + "." + base64URLEncode(signature)

	return jwt, nil

}

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
