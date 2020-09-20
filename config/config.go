package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//AppConfig is the application configuration schema
type AppConfig struct {
	MongoConfig *MongoConfig
}

// MongoConfig describes a mongo cluster configuration
type MongoConfig struct {
	MongoURL          string
	MongoUser         string
	MongoPassword     string
	MongoClusterAddr1 string
	MongoClusterAddr2 string
	MongoClusterAddr3 string
}

func createConfig() *AppConfig {
	var Config *AppConfig

	err := godotenv.Load()
	if err != nil {
		log.Printf("%v", errors.New("Could not load .env"))
	}

	Config = &AppConfig{
		MongoConfig: &MongoConfig{
			MongoURL:          *getEnv("MONGO_URL"),
			MongoUser:         *getEnv("MONGO_USER"),
			MongoPassword:     *getEnv("MONGO_PASSWORD"),
			MongoClusterAddr1: *getEnv("MONGO_CLUSTER_ADDR_1"),
			MongoClusterAddr2: *getEnv("MONGO_CLUSTER_ADDR_2"),
			MongoClusterAddr3: *getEnv("MONGO_CLUSTER_ADDR_3"),
		},
	}

	return Config
}

func getEnv(varName string) *string {
	if varVal, exists := os.LookupEnv(varName); exists {
		fmt.Printf("\nLoaded env variable %v: %v", varName, varVal)
		return &varVal
	}
	log.Panicf("\nCould not find env variable: %v", varName)
	return nil
}

// Config instance of AppConfig
var Config = createConfig()
