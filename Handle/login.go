package Handle

import (
	"fmt"
	"net/http"
	"os"
)

var (
	link string
)

func Login(w http.ResponseWriter, r *http.Request) {
	if CheckData() == false {
		http.Redirect(w, r, "/connect/install", http.StatusTemporaryRedirect)
	}

	switch os.Getenv("connection") {
	case "DEV":
		link = fmt.Sprintf("https://auth-dev.vatsim.net/oauth/authorize?client_id=%v&redirect_uri=%v&response_type=code&scope=%v&state=fdfdadsfasfsadfsfdsafds", os.Getenv("client_id"), os.Getenv("redirect"), os.Getenv("scopes"))
	case "LIVE":
		link = fmt.Sprintf("https://auth.vatsim.net/oauth/authorize?client_id=%v&redirect_uri=%v&response_type=code&scope=%v&state=fdfdadsfasfsadfsfdsafds", os.Getenv("client_id"), os.Getenv("redirect"), os.Getenv("scopes"))
	default:
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}

	http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}