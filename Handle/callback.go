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
	Token string `json:"access_token"`
}

type User struct {
	Data Data `json:"data"`
}

type Data struct {
	CID int `json:"cid"`
}

func Callback(w http.ResponseWriter, r *http.Request) {
	code := ParseResponse(w, r)
	access := AccessToken(code, w, r)
	GetData(access.Token, w, r)
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

	body, errorReading := ioutil.ReadAll(request.Body)
	if errorReading != nil {
		log.Fatal(errorReading)
	}

	var res Access
	errDecoding := json.Unmarshal(body, &res)

	if errDecoding != nil {
		log.Fatal(errDecoding)
	}

	return res
}

func GetData(access_token string, w http.ResponseWriter, r *http.Request) {
	request, err := http.NewRequest("GET", "https://auth.vatsim.net/api/user", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Bearer", access_token)
	request.Header.Add("accept", "application/json")
	client := http.Client{}
	client.Do(request)

	defer request.Body.Close()

	body, errReading := ioutil.ReadAll(request.Body)
	if errReading != nil {
		log.Fatal(errReading)
	}


	var userDetails map[string]interface{}
	errJSON := json.Unmarshal(body, &userDetails)
	if errJSON != nil {
		log.Fatal(errJSON)
	}
	fmt.Println(userDetails)
}

