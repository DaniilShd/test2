package repository

import "github.com/DaniilShd/test2/models"

type DatabaseRepo interface {
	GetAllPersons() (*[]models.Person, error)
	GetPersonByID(id int) (*models.Person, error)
	GetPersonsByOffset(filter *models.Filter) (*[]models.Person, error)
	DeletePersonByID(id int) error
	UpdatePersonByID(person *models.Person) error
	InsertPerson(person *models.Person) error
}
