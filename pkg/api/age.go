package api

import (
	"encoding/json"
	"fmt"

	"github.com/DaniilShd/test2/models"
	"github.com/gofiber/fiber/v2"
)

func GetAge(c *fiber.Ctx, name string, person *models.Person) (err error) {
	agent := fiber.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	err = json.Unmarshal(body, &person)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	return nil

}
