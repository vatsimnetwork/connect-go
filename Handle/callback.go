package Handle

// Connect callback after successful client authentication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// An Access represents access token which we get from AccessToken function
type Access struct {
	Token string `json:"access_token"`
}

// An User represents response from Connect which we get in GetData function
type User struct {
	Data Data `json:"data"`
}

// Data represents user details return from Connect in GetData function
type Data struct {
	CID int `json:"cid"`
}

// Callback returns user details
func Callback(w http.ResponseWriter, r *http.Request) {

	code := ParseResponse(w, r)

	if code != "" {

		access, err := AccessToken(code, w, r)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Fprintf(w, access.Token)
		GetData(access.Token, w, r)
	}
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
func GetData(accessToken string, w http.ResponseWriter, r *http.Request) {

	request, err := http.NewRequest("GET", "https://auth.vatsim.net/api/user", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	request.Header.Add("Bearer", accessToken)
	request.Header.Add("accept", "application/json")
	client := http.Client{}
	resp, clientErr := client.Do(request)

	if clientErr != nil {
		log.Fatal(clientErr)
		return
	}

	defer resp.Body.Close()

	body, errReading := ioutil.ReadAll(resp.Body)
	if errReading != nil {
		log.Fatal(errReading)
		return
	}


	var userDetails map[string]interface{}
	errJSON := json.Unmarshal(body, &userDetails)
	if errJSON != nil {
		log.Fatal(errJSON)
		return
	}

	fmt.Println(userDetails)
}

