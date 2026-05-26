package cmd

import (
	"sample-server/config"
	"sample-server/infra/db"
	"sample-server/repo"
	"sample-server/rest"
	"sample-server/rest/handlers/products"
	"sample-server/rest/handlers/users"
	"sample-server/rest/middlewares"
)

func Serve() {
	cfg := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(cfg)

	dbCon, err := db.NewConnection()
	if err != nil {
		panic(err)
	}

	userRepo := repo.NewUserRepo()
	userHandler := users.NewHandler(userRepo, cfg)

	productRepo := repo.NewProductRepo()
	productHandler := products.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cfg, userHandler, productHandler)
	server.Start()
}
