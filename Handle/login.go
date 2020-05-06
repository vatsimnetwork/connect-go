package Handle

import (
	"connect-go/layout"
	"fmt"
	"net/http"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if CheckData() == false {
		http.Redirect(w, r, "/install", http.StatusTemporaryRedirect)
	}

	url := "https://auth.vatsim.net/oauth/authorize?client_id=" + os.Getenv("client_id") + "&redirect_uri" + os.Getenv("redirect") + "&response_type=code&scope=full_name vatsim_details&state=fdfdadsfasfsadfsfdsafds"

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callback() {

}

func Env(w http.ResponseWriter, r *http.Request) {
	if CheckData() == true {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
	if r.Method == "GET" {
		form := layout.Layout("", "<form method=\"POST\" action=\"\">" +
			"<center><input class=\"form-control w-50 m-2\" name=\"client_id\" placeholder=\"Please enter client id\"></center>" +
			"<center><input class=\"form-control w-50 m-2\" name=\"redirect\" placeholder=\"Please enter redirect URL\"></center>" +
			"<center><input class=\"form-control w-50 m-2\" name=\"secret\" placeholder=\"Please enter client secret\"></center>" +
			"<center><button type=\"submit\" class=\"btn btn-success\">Save</button></center>" +
			"</form>", "")
		fmt.Fprintf(w, form)
	} else {
		r.ParseForm()
		os.Setenv("client_id", r.FormValue("client_id"))
		os.Setenv("redirect", r.FormValue("redirect"))
		os.Setenv("secret", r.FormValue("secret"))

		http.Redirect(w, r, "/", http.StatusCreated)
	}
}