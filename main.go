package connect_go

import (
	"connect-go/handle"
	"net/http"
)

func Url() {
	http.HandleFunc("/connect/login", Handle.Login)
	http.HandleFunc("/connect/install", Handle.Env)
	http.HandleFunc("/connect/callback", Handle.Callback)
}