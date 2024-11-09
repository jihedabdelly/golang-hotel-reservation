package api

import (
	"golang-hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

// this hnadler is public (upperCase 1st letter)
func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Walter",
		LastName: "in the meth lab",
	}
	return c.JSON(u)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("James")
}