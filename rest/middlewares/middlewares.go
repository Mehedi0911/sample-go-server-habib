package middlewares

import "sample-server/config"

type Middlewares struct {
	cfg *config.Config
}

func NewMiddlewares(cfg *config.Config) *Middlewares {
	return &Middlewares{cfg: cfg}
}
