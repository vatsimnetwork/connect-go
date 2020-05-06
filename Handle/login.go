package Handle

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://matanbudimir.de", http.StatusPermanentRedirect)
}

func callback() {

}
