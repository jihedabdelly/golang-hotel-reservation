package api

import "github.com/gofiber/fiber/v2"

// this hnadler is public (upperCase 1st letter)
func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("James")
}