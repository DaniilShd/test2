package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/DaniilShd/test2/pkg/api"
	"github.com/DaniilShd/test2/pkg/api/filter"
	"github.com/DaniilShd/test2/pkg/api/pagination"

	"github.com/DaniilShd/test2/internal/config"
	"github.com/DaniilShd/test2/internal/driver"
	"github.com/DaniilShd/test2/internal/repository"
	"github.com/DaniilShd/test2/internal/repository/dbrepo"
	"github.com/DaniilShd/test2/models"
	"github.com/gofiber/fiber/v2"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepository(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) GetPersons(c *fiber.Ctx) error {

	filter, err := filter.GetFilter(c)
	if err != nil {
		fmt.Println(err)
	}

	err = pagination.GetPagination(c, filter)
	if err != nil {
		fmt.Println(err)
	}

	persons, err := m.DB.GetPersonsByOffset(filter)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(persons)

	return c.Render("index", fiber.Map{
		"Flter":   filter,
		"Persons": persons,
	})
}

func (m *Repository) DeletePersonByID(c *fiber.Ctx) error {

	id := c.Params("id")

	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	err = m.DB.DeletePersonByID(ID)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (m *Repository) ChangePersonByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var person models.Person
	ID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	person.ID = ID

	err = json.NewDecoder(bytes.NewReader(c.Body())).Decode(&person)
	if err != nil {
		fmt.Println(err)
	}

	err = m.DB.UpdatePersonByID(&person)
	if err != nil {
		fmt.Println(err)
	}

	return c.Render("index", fiber.Map{
		"Title": "Пут",
	})
}

func (m *Repository) AddPerson(c *fiber.Ctx) error {

	var person models.Person

	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&person)
	if err != nil {
		fmt.Println(err)
	}

	err = api.GetAge(c, person.Name, &person)
	if err != nil {
		fmt.Println(err)
	}

	err = api.GetGender(c, person.Name, &person)
	if err != nil {
		fmt.Println(err)
	}

	err = api.GetNationality(c, person.Name, &person)
	if err != nil {
		fmt.Println(err)
	}

	err = m.DB.InsertPerson(&person)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
