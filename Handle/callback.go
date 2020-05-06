package Handle

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"net/url"
)
var (
	link string
)
func Callback(w http.ResponseWriter, r *http.Request) {
	code := ParseResponse(w, r)
	AccessToken(code, w, r)
}

func ParseResponse(w http.ResponseWriter,r *http.Request) string {
	parseError := r.ParseForm()
	if parseError != nil {
		log.Fatal(parseError)
	}
	code := r.FormValue("code")
	return code
}

func AccessToken(code string, w http.ResponseWriter,r *http.Request) {
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
	data.Add("client_id", os.Getenv("client_id"))
	data.Add("client_secret", os.Getenv("secret"))
	data.Add("redirect_uri", os.Getenv("redirect"))
	data.Add("code", code)

	request, requestError := http.PostForm(link, data)

	if requestError != nil {
		log.Fatal(requestError)
	}

	fmt.Println(request)

}

