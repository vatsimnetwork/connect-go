package handle

import (
	"connect-go/helpers"
	"fmt"
	"net/http"
	"os"
)

// Link variable (package wide)
var (
	link string
)

// Login function sends a GET request to connect with entered credentials
func Login(w http.ResponseWriter, r *http.Request) {
	if helpers.CheckData() == false {
		http.Redirect(w, r, "/connect/install", http.StatusTemporaryRedirect)
	}

	switch os.Getenv("connection") {
	case "DEV":
		link = fmt.Sprintf("https://auth-dev.vatsim.net/oauth/authorize?client_id=%v&redirect_uri=%v&response_type=code&scope=%v&state=%v", os.Getenv("client_id"), os.Getenv("redirect"), os.Getenv("scopes"), helpers.RandString(40))
	case "LIVE":
		link = fmt.Sprintf("https://auth.vatsim.net/oauth/authorize?client_id=%v&redirect_uri=%v&response_type=code&scope=%v&state=%v", os.Getenv("client_id"), os.Getenv("redirect"), os.Getenv("scopes"), helpers.RandString(40))
	default:
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}

	http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}