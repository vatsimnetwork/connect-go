package main

import (
	"connect-go/handle"
	index "connect-go/layout"
	"fmt"
	"net/http"
	"os"
)
func handleMain(w http.ResponseWriter, r *http.Request) {
	if Handle.CheckData() == false {
		http.Redirect(w, r, "/install", http.StatusTemporaryRedirect)
	}

	var html = index.Layout("", "<div style=\"margin-top: 200px;\">" +
		"<center>" +
		"<h2 style=\"color: gray;\">VATSIM Connect Authentication</h2><button class=\"btn btn-success\" onclick=\"location.href='login'\">Login</a>" +
		"</center>" +
		"</div>", "")
	fmt.Fprintf(w, html)
}

func main() {
		http.HandleFunc("/", handleMain)
		http.HandleFunc("/login", Handle.Login)
		http.HandleFunc("/install", Handle.Env)
		http.ListenAndServe(":69", nil)
		os.Unsetenv("client_id")
		os.Unsetenv("redirect")
		os.Unsetenv("secret")
}