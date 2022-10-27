package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(routes *mux.Router, service *Service) {
	routes.Handle("/", hello())
	routes.Handle("/contacts", getContacts(service)).
		Methods(http.MethodGet)
	routes.Handle("/contacts/{id}", getContact(service)).
		Methods(http.MethodGet)
	routes.Handle("/contacts", createContacts(service)).
		Methods(http.MethodPost)
}

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

func getContacts(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contacts := service.GetContacts()
		response, err := json.Marshal(contacts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		w.Write(response)
	})
}

func createContacts(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()

		c := Contact{}

		err := json.NewDecoder(body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		service.CreateContact(c)
		fmt.Fprintf(w, "Contact created")
	})
}

func getContact(service *Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		contact := service.GetContact(id)
		response, err := json.Marshal(contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		w.Write(response)
	})
}
