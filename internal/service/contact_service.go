package service

import (
	"time"

	"github.com/esvarez/go-course/internal/model"
)

type contactsDB interface {
	GetContacts() []model.Contact
	CreateContact(contact model.Contact)
	GetContact(id string) (model.Contact, error)
	UpdateContact(id string, contactUpdates model.Contact)
	DeleteContact(id string)
}

type Service struct {
	//db map[int]model.Contact\
	db contactsDB
}

func NewService(db contactsDB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetContacts() []model.Contact {
	return s.db.GetContacts()
	// contacts := make([]model.Contact, 0)
	// for _, contact := range s.db {
	// 	contacts = append(contacts, contact)
	// }
	// return contacts
}

func (s *Service) CreateContact(contact model.Contact) {
	contact.ID = int(time.Now().Unix())
	//s.db[contact.ID] = contact
	s.db.CreateContact(contact)
}

func (s *Service) GetContact(id string) (model.Contact, error) {
	// idint, _ := strconv.Atoi(id)
	// contact, exist := s.db[idint]

	// if !exist {
	// 	return model.Contact{}, errors.New("Contact not found")
	// }

	// return contact, nil
	return s.db.GetContact(id)
}

func (s *Service) UpdateContact(id string, contactUpdates model.Contact) {
	// idint, _ := strconv.Atoi(id)
	// contactUpdates.ID = idint
	// s.db[idint] = contactUpdates
	s.db.UpdateContact(id, contactUpdates)
}

func (s *Service) DeleteContact(id string) {
	// idint, _ := strconv.Atoi(id)
	// delete(s.db, idint)
	s.db.DeleteContact(id)
}
