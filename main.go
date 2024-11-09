package main

import (

	
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()
	app.Get("/foo", handleFoo)
	app.Listen(":5500")
}

func handleFoo(c *fiber.Ctx ) error {

	return c.JSON(map[string]string{"msg": "endpoint reached vol 3"})
}