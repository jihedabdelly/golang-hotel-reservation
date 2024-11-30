package db

const MongoDBNameEnvName = "MONGO_DB_NAME"

const (
	DBNAME      = "hotel-reservation"
	DBNAME_TEST = "hotel-reservation-test"
	DBURI       = "mongodb://localhost:27017"
)

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
