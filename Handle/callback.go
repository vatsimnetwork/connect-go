package Handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Access struct {
	token string `json:"access_token"`
}

func Callback(w http.ResponseWriter, r *http.Request) {
	code := ParseResponse(w, r)
	accessToken := AccessToken(code, w, r)
	fmt.Fprintf(w, accessToken.token)
}

func ParseResponse(w http.ResponseWriter,r *http.Request) string {
	parseError := r.ParseForm()
	if parseError != nil {
		log.Fatal(parseError)
	}
	code := r.FormValue("code")
	return code
}

func AccessToken(code string, w http.ResponseWriter,r *http.Request) Access {
	switch os.Getenv("connection") {
	case "DEV":
		link = "https://auth-dev.vatsim.net/oauth/token"
	case "LIVE":
		link = "https://auth.vatsim.net/oauth/token"
	default:
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}


	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", os.Getenv("client_id"))
	data.Set("client_secret", os.Getenv("secret"))
	data.Set("redirect_uri", os.Getenv("redirect"))
	data.Set("code", code)

	request, requestError := http.PostForm(link, data)
	request.Header.Set("Content-Type", "application/json")

	if requestError != nil {
		log.Fatal(requestError)
	}


	defer request.Body.Close()

	read, errorReading := ioutil.ReadAll(request.Body)
	if errorReading != nil {
		log.Fatal(errorReading)
	}

	var res Access
	errDecoding := json.Unmarshal(read, &res)

	if errDecoding != nil {
		log.Fatal(errDecoding)
	}

	return res

}

