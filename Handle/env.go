package Handle

import (
	"fmt"
	"os"
	"net/http"
	"connect-go/layout"
)

func Env(w http.ResponseWriter, r *http.Request) {
	if CheckData() == true {
		error := layout.Layout("", "<form method=\"POST\" action=\"\">" +
			"<center><h1 class=\"t-red\">Error Occurred</h1></center>" +
			"</form>", "")
		fmt.Fprintf(w, error)
	}
	if r.Method == "GET" {
		form := layout.Layout("", "<form method=\"POST\" action=\"\">" +
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