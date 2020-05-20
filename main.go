package main

import (
	"log"
	"net/http"
	"os"
	"si-convert/api"
)

func main() {
	port := 8082
	logger := log.New(os.Stdout, "[SI-Convert] ", log.LstdFlags|log.Lshortfile)
	a := api.New(logger)
	logger.Printf("Api is listening on port: %v", port)
	if err := a.Start(port); err != http.ErrServerClosed {
		logger.Fatalf("Api has closed unexpectly: %v", err)
	}
	logger.Print("Exiting Api service")
}