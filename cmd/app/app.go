package main

import (
	"path/filepath"

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
	os.Setenv("CLOUD_HOME", filepath.Join(os.Getenv("CLOUD_HOME"), "companycloud/"))

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.AddHook(&securefilechanger.FileLogHook{})
	logrus.SetReportCaller(true)

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error init enviroments: %v", err.Error())
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     dbPort,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   "postgres",
	})

	if err != nil {
		logrus.Fatalf("error init Database: %v", err.Error())
	}

	logrus.Info("[app file storage] ", os.Getenv("CLOUD_HOME"))

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := api.NewHandler(services)

	server := securefilechanger.NewServer()
	go func() {
		err := server.Run("8080", handlers.InitRouter())
		if err != nil {
			logrus.Fatalf("error starting server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-sigChan
	logrus.Warnf("stop app with sys call: %v", sig)

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error while shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error while db close: %s", err.Error())
	}
}
