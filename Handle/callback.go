package Handle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	code := ParseResponse(w, r)
	AccessToken(code)
}

func ParseResponse(w http.ResponseWriter,r *http.Request) string {
	parseError := r.ParseForm()
	if parseError != nil {
		fmt.Fprintf(w,"could not parse query: %v", parseError)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")
	return code
}

func AccessToken(code string) string {
	
}
