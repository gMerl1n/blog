package server

import (
	"net/http"

	"github.com/gMerl1n/blog/internal/config"
	"github.com/gMerl1n/blog/internal/handlers"
)

func NewServer(cfg *config.ConfigServer, handler handlers.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler.InitRoutes().Handler(),
	}

}
