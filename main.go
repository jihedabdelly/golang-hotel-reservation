package main

import (

	
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiV1.Get("/user", handleUser)
	
	app.Listen(":5500")
}

func handleFoo(c *fiber.Ctx ) error {
	return c.JSON(map[string]string{"msg": "result from /foo endpoint"})
}

func handleUser(c *fiber.Ctx ) error {
	return c.JSON(map[string]string{"user": "walter white"})
}