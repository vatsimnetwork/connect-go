package Handle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")
	fmt.Println(code)
	ParseResponse(r)
}

func ParseResponse(r *http.Request) {
	defer r.Body.Close()
	read, errReading := ioutil.ReadAll(r.Body)

	if errReading != nil {
		log.Fatal(errReading)
	}

	//var response map[string]interface{}

	//errDecode := json.Unmarshal(read, &response)

	/*if errDecode != nil {
		log.Fatal(errDecode)
	}*/

	fmt.Println(read)
}
