package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esvarez/go-course/internal/model"
	"github.com/gorilla/mux"
)

type contactsService interface {
	GetContacts() []model.Contact
	GetContact(id string) (model.Contact, error)
	CreateContact(contact model.Contact)
	UpdateContact(id string, contact model.Contact)
	DeleteContact(id string)
}

type ContactController struct {
	contactSrv contactsService
}

func NewContactController(s contactsService) ContactController {
	return ContactController{
		contactSrv: s,
	}
}

func NewHandler(routes *mux.Router, ctrl ContactController) {
	routes.Handle("/", hello())
	routes.Handle("/contacts", ctrl.getContacts()).
		Methods(http.MethodGet)
	routes.Handle("/contacts/{id}", ctrl.getContact()).
		Methods(http.MethodGet)
	routes.Handle("/contacts", ctrl.createContacts()).
		Methods(http.MethodPost)
	routes.Handle("/contacts/{id}", ctrl.updateContact()).
		Methods(http.MethodPut)
	routes.Handle("/contacts/{id}", ctrl.deleteContact()).
		Methods(http.MethodDelete)
}

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

func (c *ContactController) getContacts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contacts := c.contactSrv.GetContacts()
		response, err := json.Marshal(contacts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		w.Write(response)
	})
}

func (c *ContactController) createContacts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()

		co := model.Contact{}

		err := json.NewDecoder(body).Decode(&co)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		c.contactSrv.CreateContact(co)
		fmt.Fprintf(w, "Contact created")
	})
}

func (c *ContactController) getContact() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		contact, err := c.contactSrv.GetContact(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		response, err := json.Marshal(contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		w.Write(response)
	})
}

func (c *ContactController) updateContact() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		body := r.Body
		defer body.Close()

		conctactUpdates := model.Contact{}

		err := json.NewDecoder(body).Decode(&conctactUpdates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		c.contactSrv.UpdateContact(id, conctactUpdates)
		fmt.Fprintf(w, "Contact updated")
	})
}

func (c *ContactController) deleteContact() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		c.contactSrv.DeleteContact(id)
		fmt.Fprintf(w, "Contact deleted")
	})
}
