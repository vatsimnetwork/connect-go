package handle

// Connect callback after successful client authentication

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"context"
)

// Callback returns user details
func Callback(w http.ResponseWriter, r *http.Request) (*User, error) {

	code := ParseResponse(w, r)

	if code != "" {

		access, err := AccessToken(code, w, r)
		if err != nil {
			log.Fatal(err)
			return &User{}, err
		}

		user, err := GetData(access.Token, w, r)
		return user, nil
	}
	log.Fatal("No code received")
	return &User{}, nil
}

// ParseResponse returns code provided by Connect
func ParseResponse(w http.ResponseWriter,r *http.Request) string {
	parseError := r.ParseForm()
	if parseError != nil {
		log.Fatal(parseError)
		return ""
	}
	code := r.FormValue("code")
	return code
}

// AccessToken function sends a request with provided code and returns access_token provided by Connect
func AccessToken(code string, w http.ResponseWriter,r *http.Request) (*Access, error) {

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
		return nil, requestError
	}

	defer request.Body.Close()

	body, errorReading := ioutil.ReadAll(request.Body)

	if errorReading != nil {
		log.Fatal(errorReading)
		return nil, errorReading
	}

	var res Access
	errDecoding := json.Unmarshal(body, &res)

	if errDecoding != nil {
		log.Fatal(errDecoding)
	}

	return &res, nil
}

// GetData function returns users details (TBD)
func GetData(accessToken string, w http.ResponseWriter, r *http.Request) (*User, error) {

	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	}))

	switch os.Getenv("connection") {
	case "LIVE":
		link = "https://auth.vatsim.net/api/user"
	case "DEV":
		link = "https://auth-dev.vatsim.net/api/user"
	default:
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

	req, err := client.Get(link)

	if err != nil {
		log.Fatal(err)
		return &User{}, err
	}

	defer req.Body.Close()

	body, errRead := ioutil.ReadAll(req.Body)

	if errRead != nil {
		log.Fatal(errRead)
		return &User{}, errRead
	}

	var user User
	errDecoding := json.Unmarshal(body, &user)
	if errDecoding != nil {
		log.Fatal(errDecoding)
		return &User{}, errDecoding
	}

	return &user, nil
}

