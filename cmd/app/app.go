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
		logrus.Fatalf("Error Init enviroments: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   "postgres",
	})

	if err != nil {
		logrus.Fatalf("Error Init Database: %v", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := api.NewHandler(services)

	server := securefilechanger.NewServer()
	go func() {
		err := server.Run("8080", handlers.InitRouter())
		if err != nil {
			logrus.Fatalf("Error Starting server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	logrus.Warnf("Stop app with sys call: %v", sig)

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("Error while shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error while db close: %s", err.Error())
	}
}
