package main

import (
	"fmt"
	"net/http"
	"connect-go/handle"
	index "connect-go/layout"
)
func handleMain(w http.ResponseWriter, r *http.Request) {
	var html = index.Layout("", "<div style=\"margin-top: 200px;\"><center><h2 style=\"color: gray;\">VATSIM Connect Authentication</h2><button class=\"btn btn-success\" onclick=\"location.href='login'\">Login</a></center></div>", "")
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", Handle.Login)
	http.ListenAndServe(":69", nil)
}