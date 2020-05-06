package Handle

import (

	"os"
)

func CheckData() (r bool) {
	if os.Getenv("client_id") == "" && os.Getenv("redirect") == "" && os.Getenv("secret") == "" {
		return false
	} else {
		return true
	}
}
