package filter

import (
	"github.com/DaniilShd/test2/models"
	"github.com/gofiber/fiber/v2"
)

//http://localhost:3000/?name=Dan&surname=Shaid&patronymic=Ser&nationality=RU&gender=male&age=25&offset=0&limit=5

func GetFilter(c *fiber.Ctx) (*models.Filter, error) {
	var filter models.Filter

	name := c.Query("name")
	if name != "" {
		filter.Name = name
		filter.Set += 1
	}

	surname := c.Query("surname")
	if surname != "" {
		filter.Surname = surname
		filter.Set += 1
	}

	patronymic := c.Query("patronymic")
	if patronymic != "" {
		filter.Patronymic = patronymic
		filter.Set += 1
	}

	nationality := c.Query("nationality")
	if nationality != "" {
		filter.Nationality = nationality
		filter.Set += 1
	}

	gender := c.Query("gender")
	if gender != "" {
		filter.Gender = gender
		filter.Set += 1
	}

	age := c.Query("age")
	if age != "" {
		filter.Age = age
		filter.Set += 1
	}

	return &filter, nil
}
