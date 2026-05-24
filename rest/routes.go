package rest

import (
	"net/http"
	"sample-server/rest/handlers"
	"sample-server/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.HandleGetProducts),
		),
	)
	// mux.Handle("GET /products", middlewares.Logger(http.HandlerFunc(handlers.HandleGetProducts)))
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.HandleCreateProducts),
		),
	)
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
}
