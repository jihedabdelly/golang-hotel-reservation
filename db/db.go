package db

import "os"

var DBNAME string

const MongoDBNameEnvName = "MONGO_DB_NAME"


type Pagination struct {
	Page  int64
	Limit int64
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}

func init() {
	DBNAME = os.Getenv("MONGO_DB_NAME")
}
