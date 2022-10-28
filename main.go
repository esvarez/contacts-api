package main

import (
	"github.com/esvarez/go-course/internal/handler"
	"github.com/esvarez/go-course/internal/service"
	"github.com/esvarez/go-course/pkg/server"
	"github.com/esvarez/go-course/pkg/storage"

	"github.com/gorilla/mux"
)

func main() {
	routes := mux.NewRouter()
	keyValueDB := storage.NewKeyvDB()

	service := service.NewService(keyValueDB)
	controller := handler.NewContactController(service)

	handler.NewHandler(routes, controller)

	server.Start(routes)
}
