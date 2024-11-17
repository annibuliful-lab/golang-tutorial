package hello

import "net/http"

func Router(mux *http.ServeMux) {
	mux.HandleFunc("GET /hello", GetHelloController)
}
