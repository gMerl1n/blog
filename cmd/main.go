package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gMerl1n/blog/internal/config"
	"github.com/gMerl1n/blog/internal/handlers"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/gMerl1n/blog/internal/services"
	"github.com/gMerl1n/blog/pkg/db"
	"github.com/gMerl1n/blog/pkg/logging"
	"github.com/gMerl1n/blog/server"
	"github.com/joho/godotenv"
)

func main() {

	context := context.Background()

	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load envs")
	}

	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	logger := logging.InitLogger(config.ConfigServer)

	db, err := db.NewPostgresDB(context, config.ConfigDB)
	if err != nil {

	}

	repos := repository.NewRepository(db, logger)
	services := services.NewService(repos, logger)
	handlers := handlers.NewHandler(services, logger)

	srv := server.NewServer(config.ConfigServer, handlers)

	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}

	}()

	logger.Print("BlogApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Print("BlogApp Shutting Down")

	if err := srv.Shutdown(context); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
