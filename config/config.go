package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var hello string
var mongoURL string
var mongoUser string
var mongoPass string

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("%v", errors.New("Could not load .env"))
	}
	hello = os.Getenv("HELLO")
	fmt.Printf("%v", hello)
}
