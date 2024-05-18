package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"context"
	"os"
	"os/signal"
	"syscall"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/api"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
	"github.com/KodokuOdius/SecureFileChanger/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error init enviroments: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   "postgres",
	})

	if err != nil {
		logrus.Fatalf("Error database: %v", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := api.NewHandler(services)

	server := new(securefilechanger.Server)
	go func() {
		logrus.Printf("---- SERVER STARTING ON PORT: %s ----\n", "8080")
		err := server.Run("8080", handlers.InitRouter())
		if err != nil {
			logrus.Printf("ERROR STARTING SERVER: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logrus.Printf("Closing now, We've gotten signal: %v", sig)
	ctx := context.Background()
	server.ShutDown(ctx)
}
