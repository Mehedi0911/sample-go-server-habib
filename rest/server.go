package rest

import (
	"fmt"
	"net/http"
	"sample-server/config"
	"sample-server/rest/middlewares"
	"strconv"
)

func Start(cfg config.Config) {
	manager := middlewares.NewManager()
	manager.Use(middlewares.InitPrinter, middlewares.CorsWithPreflight, middlewares.Logger)
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	wrappedMux := manager.With(mux)
	// routerHandler := middlewares.CorsWithPreflight(mux)
	fmt.Println("Server is running on port", cfg.HTTPPort)
	addr := ":" + strconv.Itoa(cfg.HTTPPort)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		panic(err)
	}
}
