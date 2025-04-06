package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gojeksrepo/config"
	"gojeksrepo/internal/router"
	"log"
	"os"
)

func main() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	// Init DB
	config.DatabaseInit()

	// Call Router
	echoCall := router.InitRouter()

	// Listen port
	portRunner := os.Getenv("PORT_RUNNER")
	echoCall.Logger.Fatal(echoCall.Start(portRunner))
}
