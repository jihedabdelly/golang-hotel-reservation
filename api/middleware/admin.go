package middleware

import (
	"fmt"
	"golang-hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return fmt.Errorf("not authorized nnn admin")
	}
	if !user.IsAdmin {
		return fmt.Errorf("not authorized not admin" )
	}

	return c.Next()
}