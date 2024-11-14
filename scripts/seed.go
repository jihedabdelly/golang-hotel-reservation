package main

import (
	"fmt"
	"golang-hotel-reservation/types"
)

func main ()  {
	
	hotel := types.Hotel{
		Name: "Bellucia",
		Location: "Maldives",
	}

	room := types.Room{
		Type: types.SingleRoomType,
		BasePrice: 99.9,
		HotelID: hotel.ID,
	}
	fmt.Println("seeding ...", room)
}