package middlewares

import "net/http"

func InitPrinter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("This is the init printer middleware")
		next.ServeHTTP(w, r)
	})
}
