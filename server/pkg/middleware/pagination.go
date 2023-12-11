package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Pagination(c *fiber.Ctx) error {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	c.Locals("pagination_page", page)
	c.Locals("pagination_page_size", pageSize)

	return c.Next()
}
