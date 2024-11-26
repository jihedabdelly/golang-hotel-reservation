package db

const MongoDBNameEnvName = "MONGO_DB_NAME"

const (
	DBNAME      = "hotel-reservation"
	DBNAME_TEST = "hotel-reservation-test"
	DBURI       = "mongodb://localhost:27017"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
