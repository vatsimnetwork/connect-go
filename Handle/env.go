package handle

import (
	"connect-go/helpers"
	"connect-go/layout"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Env functions allow owners to enter Connect credentials and save them a environmental variables
// Checks are credentials already entered, if they are it will return "Forbidden"
// If credentials aren't entered and request is GET, function returns a form
// If request isn't get it will save sent form
func Env(w http.ResponseWriter, r *http.Request) {

	if helpers.CheckData() == true {

		error := layout.Generate("", "<form method=\"POST\" action=\"\">" +
			"<center><h1 style=\"color: red;\">Forbidden</h1></center>" +
			"</form>", "")

		fmt.Fprintf(w, error)

		ip := strings.Split(r.RemoteAddr, ":")[0]

		fmt.Println(ip)

	} else if r.Method == "GET" {

		form := Layout.Generate("", "<form method=\"POST\" action=\"\">" +
			"<center>" +
			"<div><label>Please enter client id</label><input class=\"form-control w-50 m-2\" type=\"number\" name=\"client_id\" placeholder=\"Please enter client id\"></div>" +
			"</center>" +
			"<center>" +
			"<div><label>Please choose a connection</label><select class=\"form-control w-50 m-2\" name=\"connection\">" +
			"<option value=\"DEV\">Development</option><option value=\"LIVE\">Live</option>" +
			"</select></div>" +
			"</center>" +
			"<center><label>Please enter redirect URI</label><input class=\"form-control w-50 m-2\" name=\"redirect\" placeholder=\"Please enter redirect URL\"></center>" +
			"<center><label>Please enter secret</label><input class=\"form-control w-50 m-2\" name=\"secret\" placeholder=\"Please enter client secret\"></center>" +
			"<center><label>Please enter scopes (separated with a space)</label><input class=\"form-control w-50 m-2\" name=\"scopes\" placeholder=\"Please enter scopes\"></center>" +
			"<center><button type=\"submit\" class=\"btn btn-success\">Save</button></center>" +
			"</form>", "")

		fmt.Fprintf(w, form)

	} else {

		r.ParseForm()

		os.Setenv("client_id", r.FormValue("client_id"))

		os.Setenv("redirect", r.FormValue("redirect"))

		os.Setenv("secret", r.FormValue("secret"))

		os.Setenv("scopes", r.FormValue("scopes"))

		os.Setenv("connection", r.FormValue("connection"))

		http.Redirect(w, r, "/", http.StatusCreated)
	}
}