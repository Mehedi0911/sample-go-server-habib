package products

import (
	"net/http"
	"sample-server/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.HandleGetProducts),
		),
	)
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductByID),
		),
	)
	// mux.Handle("GET /products", middlewares.Logger(http.HandlerFunc(handlers.HandleGetProducts)))
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.HandleCreateProducts),
			h.middlewares.AuthJWT,
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.HandleUpdateProducts),
			h.middlewares.AuthJWT,
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.HandleDeleteProductByID),
			h.middlewares.AuthJWT,
		),
	)

}
