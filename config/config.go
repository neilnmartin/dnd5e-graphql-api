package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	mongoURL      string
	mongoUser     string
	mongoPassword string
}

// Anything test
var Anything = "anything"

func new() *appConfig {
	var Config appConfig

	err := godotenv.Load()
	if err != nil {
		log.Printf("%v", errors.New("Could not load .env"))
	}

	Config = appConfig{
		mongoURL:      *getEnv("MONGO_URL"),
		mongoUser:     *getEnv("MONGO_USER"),
		mongoPassword: *getEnv("MONGO_PASSWORD"),
	}

	return &Config
}

func getEnv(varName string) *string {
	if varVal, exists := os.LookupEnv("MONGO_URL"); exists {
		fmt.Printf("Loaded env variable %v: %v", varName, varVal)
		return &varVal
	}
	log.Panicf("Could not find env variable: %v", varName)
	return nil
}

// Config holds the application config variables
var Config = new()
