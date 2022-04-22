package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Database string connection
	DatabaseStringConnection = ""
	// Port API run on this port
	Port = 0

	// SecretKey is the key to sign jwt
	SecretKey []byte
)

// Load will initialize env vars
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	DatabaseStringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
