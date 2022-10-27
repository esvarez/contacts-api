package main

import "github.com/gorilla/mux"

func main() {
	routes := mux.NewRouter()
	db := make(map[int]Contact)

	service := NewService(db)

	NewHandler(routes, service)

	Start(routes)
}
