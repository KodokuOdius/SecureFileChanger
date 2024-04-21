package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/KodokuOdius/SecureFileChanger/api/v1"
	"github.com/gorilla/mux"
)

const (
	port = "8080"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/upload", v1.HandleUpload).Methods("POST")
	s.HandleFunc("/download", v1.HandleDownload).Methods("GET")
	s.HandleFunc("/user/create", v1.HandleCreateUser).Methods("POST")

	server := http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     s,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		log.Printf("---- SERVER STARTING ON PORT: %s ----\n", port)
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("ERROR STARTING SERVER: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("Closing now, We've gotten signal: %v", sig)
	ctx := context.Background()
	server.Shutdown(ctx)
}
