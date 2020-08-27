package helpers

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

var (
	client_id, redirect, secret, scopes, connection string
)

// load function loads the environmental variables
func load() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	client_id = os.Getenv("CLIENT_ID")
	redirect = os.Getenv("REDIRECT")
	scopes = os.Getenv("SCOPES")
	secret = os.Getenv("SECRET")
	connection = os.Getenv("CONNECTION")

}

// CheckData function checks are any of these environmental variables empty
func CheckData() (r bool) {
	load()
	if client_id == "" || redirect == "" || secret == "" || scopes == "" || connection == "" {
		return false
	}
	return true
}
