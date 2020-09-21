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
	Persistence string
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

	err := godotenv.Load()
	if err != nil {
		log.Printf("%v", errors.New("Could not load .env"))
	}

	config := &AppConfig{
		Persistence: *getVar("PERSISTENCE"),
		MongoConfig: &MongoConfig{
			MongoClusterAddr1: *getVar("MONGO_CLUSTER_ADDR_1"),
			MongoClusterAddr2: *getVar("MONGO_CLUSTER_ADDR_2"),
			MongoClusterAddr3: *getVar("MONGO_CLUSTER_ADDR_3"),
			MongoUser:         *getVar("MONGO_USER"),
			MongoPassword:     *getVar("MONGO_PASSWORD"),
		},
	}

	return config
}

func getVar(varName string) *string {
	if varVal, exists := os.LookupEnv(varName); exists {
		fmt.Printf("\nLoaded env variable %v: %v", varName, varVal)
		return &varVal
	}
	log.Panicf("\nCould not find env variable: %v", varName)
	return nil
}

// Config instance of AppConfig
var Config = createConfig()
