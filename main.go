package main

import (
	"context"
	"flag"
	"golang-hotel-reservation/api"
	"golang-hotel-reservation/db"
	"log"
	

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"


var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}


func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	//handler initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))

	listenAddress := flag.String("listenAddress", ":5500", "The listen address or port of the API server.")
	flag.Parse()

	app := fiber.New(config)
	apiV1 := app.Group("/api/v1")

	apiV1.Put("/user/:id", userHandler.HandlePutUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiV1.Post("/user", userHandler.HandlePostUser)
	apiV1.Get("/user", userHandler.HandleGetUsers )
	apiV1.Get("/user/:id", userHandler.HandleGetUser )

	app.Listen(*listenAddress)
}
