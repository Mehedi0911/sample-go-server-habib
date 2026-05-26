package users

import (
	"net/http"
	"sample-server/rest/middlewares"
)

func (h *Handler) RegisterUserRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.HandleCreateUser),
		),
	)
	mux.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(h.HandleLogin),
		),
	)
}
