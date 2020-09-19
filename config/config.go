package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	Hello string
}

// Anything test
var Anything = "anything"

func new() *appConfig {
	var Config appConfig
	fmt.Println("config main is running")
	err := godotenv.Load()
	if err != nil {
		log.Printf("%v", errors.New("Could not load .env"))
	}
	if Hello, exists := os.LookupEnv("MONGO_URL"); exists {
		fmt.Println("config load env hello", Hello)
		Config = appConfig{
			Hello,
		}
		return &Config
	}
	fmt.Println("could not find hello")
	return nil
}

// Config holds the application config variables
var Config = new()
