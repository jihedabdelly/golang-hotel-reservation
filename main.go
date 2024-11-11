package main

import (
	"context"
	"flag"
	"fmt"
	"golang-hotel-reservation/api"
	"golang-hotel-reservation/types"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "James",
		LastName: "At the water cooler",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil { 
		log.Fatal(err)
	}

	var james types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err!= nil {
		log.Fatal(err)
	} 

	fmt.Println(james)

	listenAddress := flag.String("listenAddress", ":5500", "The listen address or port of the API server.")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddress)
}
