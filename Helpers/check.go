package Helpers

import (
	"os"
)

// CheckData function checks are any of these environmental variables empty
func CheckData() (r bool) {
	if os.Getenv("client_id") == "" || os.Getenv("redirect") == "" || os.Getenv("secret") == "" || os.Getenv("scopes") == "" || os.Getenv("connection") == "" {
		return false
	}
	return true
}
