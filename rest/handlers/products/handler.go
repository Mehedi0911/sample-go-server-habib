package products

import (
	"sample-server/repo"
	"sample-server/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middlewares *middlewares.Middlewares, productRepo repo.ProductRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
