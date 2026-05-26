package rest

import (
	"fmt"
	"net/http"
	"sample-server/config"
	"sample-server/rest/handlers/products"
	"sample-server/rest/handlers/users"
	"sample-server/rest/middlewares"
	"strconv"
)

type Server struct {
	cfg            *config.Config
	userHandler    *users.Handler
	productHandler *products.Handler
}

func NewServer(cfg *config.Config, userHandler *users.Handler, productHandler *products.Handler) *Server {
	return &Server{
		cfg:            cfg,
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(middlewares.InitPrinter, middlewares.CorsWithPreflight, middlewares.Logger)
	mux := http.NewServeMux()

	server.userHandler.RegisterUserRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.With(mux)
	// routerHandler := middlewares.CorsWithPreflight(mux)
	fmt.Println("Server is running on port", server.cfg.HTTPPort)
	addr := ":" + strconv.Itoa(server.cfg.HTTPPort)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		panic(err)
	}
}
