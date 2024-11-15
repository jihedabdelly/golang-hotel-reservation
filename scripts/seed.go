package main

import (
	"context"
	"fmt"
	"golang-hotel-reservation/db"
	"golang-hotel-reservation/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main ()  {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name: "Bellucia",
		Location: "Maldives",
	}

	room := types.Room{
		Type: types.SingleRoomType,
		BasePrice: 99.9,
	} 

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = insertedHotel.ID

	fmt.Println(insertedHotel)
}