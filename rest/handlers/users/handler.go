package users

import (
	"sample-server/config"
	"sample-server/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	cfg      *config.Config
}

func NewHandler(userRepo repo.UserRepo, cfg *config.Config) *Handler {
	return &Handler{
		userRepo: userRepo,
		cfg:      cfg,
	}
}
