package Handle

import (
	"net/http"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if CheckData() == false {
		http.Redirect(w, r, "/connect/install", http.StatusTemporaryRedirect)
	}

	url := "https://auth-dev.vatsim.net/oauth/authorize?client_id=" + os.Getenv("client_id") + "&redirect_uri" + os.Getenv("redirect") + "&response_type=code&scope=" + os.Getenv("scopes") + "&state=fdfdadsfasfsadfsfdsafds"

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}