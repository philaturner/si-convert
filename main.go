package main

import (
	"log"
	"net/http"
	"os"
	"si-convert/api"
)

func main() {
	port := 8080
	logger := log.New(os.Stdout, "[SI-Convert] ", log.LstdFlags|log.Lshortfile)
	a := api.New(logger)
	log.Printf("Api is listening on port: %v", port)
	if err := a.Start(port); err != http.ErrServerClosed {
		log.Fatalf("Api has closed unexpectly: %v", err)
	}
	log.Print("Exiting Api service")
}