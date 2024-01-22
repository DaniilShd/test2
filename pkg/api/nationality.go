package api

import (
	"encoding/json"
	"fmt"

	"github.com/DaniilShd/test2/models"
	"github.com/gofiber/fiber/v2"
)

func GetNationality(c *fiber.Ctx, name string, person *models.Person) (err error) {
	agent := fiber.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	var result fiber.Map
	err = json.Unmarshal(body, &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	var maxProbability float64
	maxProbability = 0
	var countryID string

	for _, s := range result {
		if rec, ok := s.([]interface{}); ok {
			for _, val := range rec {
				if count, ok := val.(map[string]interface{}); ok {
					if country_id, ok := count["country_id"].(string); ok {
						if probability, ok := count["probability"].(float64); ok {
							if probability > maxProbability {
								countryID = country_id
								maxProbability = probability
							}
						}
					}
				}
			}
		}
	}

	person.Nationality = countryID

	fmt.Println(result["country"])

	return nil

}
