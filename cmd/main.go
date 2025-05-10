package main

import (
	"gojeksrepo/config"
	"gojeksrepo/internal/router"
	"gojeksrepo/pkg"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	rootPath, err := filepath.Abs("../")
	if err != nil {
		log.Fatal(err)
	}
	env := filepath.Join(rootPath, ".env")
	errEnv := godotenv.Load(env)

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	// Init DB
	config.DatabaseInit()

	// Call Router
	echoCall := router.InitRouter()

	go pkg.KafkaReadOrderCreated()

	// Listen port
	portRunner := os.Getenv("PORT_RUNNER")
	echoCall.Logger.Fatal(echoCall.Start(portRunner))
}
