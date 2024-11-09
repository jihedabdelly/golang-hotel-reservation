package main

import (
	"flag"
	"golang-hotel-reservation/api"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	listenAddress := flag.String("listenAddress", ":5500", "The listen address or port of the API server.")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	
	app.Listen(*listenAddress)
}