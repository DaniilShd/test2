package pagination

import (
	"github.com/DaniilShd/test2/models"
	"github.com/gofiber/fiber/v2"
)

func GetPagination(c *fiber.Ctx, filter *models.Filter) error {

	offset := c.Query("offset")
	if offset != "" {
		filter.Offset = offset
	}

	limit := c.Query("limit")
	if limit != "" {
		filter.Limit = limit
	}

	return nil
}
