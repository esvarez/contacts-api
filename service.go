package main

import (
	"strconv"
	"time"
)

type Service struct {
	db map[int]Contact
}

func NewService(db map[int]Contact) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetContacts() []Contact {
	contacts := make([]Contact, 0)
	for _, contact := range s.db {
		contacts = append(contacts, contact)
	}
	return contacts
}

func (s *Service) CreateContact(contact Contact) {
	contact.ID = int(time.Now().Unix())
	s.db[contact.ID] = contact
}

func (s *Service) GetContact(id string) Contact {
	idint, _ := strconv.Atoi(id)
	return s.db[idint]
}
