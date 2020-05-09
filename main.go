package connect

import (
	"connect-go/handle"
	"net/http"
)

// URL function returns connect urls to main application function
func URL() {
	http.HandleFunc("/connect/login", Handle.Login)
	http.HandleFunc("/connect/install", Handle.Env)
}