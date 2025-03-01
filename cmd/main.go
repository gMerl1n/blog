package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gMerl1n/blog/internal/config"
	"github.com/gMerl1n/blog/internal/handlers"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/gMerl1n/blog/internal/services"
	"github.com/gMerl1n/blog/pkg/db"
	"github.com/gMerl1n/blog/pkg/jwt"
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
		fmt.Println(err)
	}

	tokenManager, err := jwt.NewManager(config.ConfigToken.JWTsecret, config.ConfigToken.OneDayInSeconds, config.ConfigToken.AccessTokenTTL, config.ConfigToken.RefreshTokenTTL)
	if err != nil {
		log.Fatal("Failed to init token manager")
		fmt.Println(err)
	}

	repos := repository.NewRepository(db, logger)
	services := services.NewService(repos, tokenManager, logger)
	handlers := handlers.NewHandler(services, tokenManager, logger)

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
