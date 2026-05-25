package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"sample-server/config"
	"sample-server/utils"
	"strings"
)

func AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//validate JWT
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessTokenParts := strings.Split(parts[1], ".")

		if len(accessTokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := accessTokenParts[0]
		jwtClaim := accessTokenParts[1]
		signature := accessTokenParts[2]

		message := jwtHeader + "." + jwtClaim

		byteArrMessage := []byte(message)

		cnf := config.GetConfig()
		byteArrSecret := []byte(cnf.JwtSecret)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := utils.Base64URLEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
