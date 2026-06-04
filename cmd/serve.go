package cmd

import (
	"fmt"
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

	dbCon, err := db.NewConnection(cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
		panic(err)
	}

	userRepo := repo.NewUserRepo(dbCon)
	userHandler := users.NewHandler(userRepo, cfg)

	productRepo := repo.NewProductRepo(dbCon)
	productHandler := products.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cfg, userHandler, productHandler)
	server.Start()
}
