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

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	//handler initialization
	var (
		hotelStore  = db.NewMongoHotelStore(client)
		roomStore   = db.NewMongoRoomStore(client, hotelStore)
		userStore   = db.NewMongoUserStore(client, db.DBNAME)
		store       = &db.Store{
			Hotel: hotelStore,
			Room:  roomStore,
			User: userStore,
		}
		hotelHandler = api.NewHotelHandler(store)
		userHandler = api.NewUserHandler(userStore)
		app          = fiber.New(config)
		apiV1        = app.Group("/api/v1")
	)

	// user handlers
	apiV1.Put("/user/:id", userHandler.HandlePutUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiV1.Post("/user", userHandler.HandlePostUser)
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)

	// hotel handlers
	apiV1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiV1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	listenAddress := flag.String("listenAddress", ":5500", "The listen address or port of the API server.")
	flag.Parse()
	app.Listen(*listenAddress)
}
