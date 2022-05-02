package main

import (
	"log"
	"os"

	"github.com/SteveYi-LAB/WHOIS-Search/internal/webserver"
	"github.com/joho/godotenv"
)

// Gin Engine
func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	webserver_listen := os.Getenv("WEBSERVER_LISTEN")

	if webserver_listen == "" {
		log.Fatal("Error to loading environment.")
	}

	webserver.Init(webserver_listen)
}
