package cmd

import (
	"sample-server/config"
	"sample-server/rest"
)

func Serve() {
	cfg := config.GetConfig()
	rest.Start(cfg)
}
