package connect

import (
	"connect-go/handle"
	"github.com/go-martini/martini"
	"net/http"
)

// URL function returns connect urls to main application function
func URL() {
	http.HandleFunc("/connect/login", Handle.Login)
	http.HandleFunc("/connect/install", Handle.Env)
}

// MartiniURL function returns url to marting application routes
func MartiniURL(m martini.ClassicMartini) {
	m.Group("/connect", func (r martini.Router) {
		r.Get("/connect/login", Handle.Login)
		r.Post("/connect/install", Handle.Env)
		r.Get("/connect/install", Handle.Env)
	})
}