package Handle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	ParseResponse(r)
}

func ParseResponse(r *http.Request) {
	defer r.Body.Close()
	read, errReading := ioutil.ReadAll(r.Body)

	if errReading != nil {
		log.Fatal(errReading)
	}

	var response map[string]interface{}

	errDecode := json.Unmarshal(read, &response)

	if errDecode != nil {
		log.Fatal(errDecode)
	}

	fmt.Println(response)
}
