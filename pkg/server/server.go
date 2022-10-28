package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newServer(port string, routes *mux.Router) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: routes,
	}
}

func Start(routes *mux.Router) {
	server := newServer(":8080", routes)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
