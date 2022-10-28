package storage

import (
	"errors"
	"strconv"

	"github.com/esvarez/go-course/internal/model"
)

type KeyvDB struct {
	db map[string]model.Contact
}

func NewKeyvDB() *KeyvDB {
	return &KeyvDB{
		db: make(map[string]model.Contact),
	}
}

func (s *KeyvDB) GetContacts() []model.Contact {
	contacts := make([]model.Contact, 0)
	for _, contact := range s.db {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (s *KeyvDB) CreateContact(contact model.Contact) {
	id := strconv.Itoa(contact.ID)
	s.db[id] = contact
}

func (s *KeyvDB) GetContact(id string) (model.Contact, error) {
	contact, exist := s.db[id]

	if !exist {
		return model.Contact{}, errors.New("Contact not found")
	}

	return contact, nil
}

func (s *KeyvDB) UpdateContact(id string, contactUpdates model.Contact) {
	// idint, _ := strconv.Atoi(id)
	// contactUpdates.ID = idint
	s.db[id] = contactUpdates
}

func (s *KeyvDB) DeleteContact(id string) {
	// idint, _ := strconv.Atoi(id)
	delete(s.db, id)
}
