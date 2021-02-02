package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConnectionDB it is the string of connection with the database
	StringConnectionDB = ""

	// Port where the API will executing
	Port = 0

	// SecretKey is the key used to sign the token
	SecretKey []byte
)

// Load load the environment variables
func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}

	StringConnectionDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
