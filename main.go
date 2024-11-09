package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	listenAddress := flag.String("listenAddress", ":5500", "The listen address or port of the API server.")
	flag.Parse()
	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiV1.Get("/user", handleUser)
	
	app.Listen(*listenAddress)
}

func handleFoo(c *fiber.Ctx ) error {
	return c.JSON(map[string]string{"msg": "result from /foo new endpoint "})
}

func handleUser(c *fiber.Ctx ) error {
	return c.JSON(map[string]string{"user": "walter white"})
}