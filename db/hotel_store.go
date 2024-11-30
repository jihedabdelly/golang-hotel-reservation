package db

import (
	"context"
	"golang-hotel-reservation/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const hotelColl = "hotels"

type HotelStore interface {
	Insert(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, GeneralizedBson, GeneralizedBson) error
	GetHotels(context.Context, GeneralizedBson, *options.FindOptions) ([]*types.Hotel, error)
	GetHotelByID(context.Context, GeneralizedBson) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll: client.Database(DBNAME).Collection(hotelColl),
	}
}

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	insertedHotel, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = insertedHotel.InsertedID.(primitive.ObjectID)

	return hotel, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, filter GeneralizedBson, update GeneralizedBson) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter GeneralizedBson, opts *options.FindOptions) ([]*types.Hotel, error) {
	resp, err := s.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var hotels []*types.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *MongoHotelStore) GetHotelByID(ctx context.Context, filter GeneralizedBson) (*types.Hotel, error) {
	var hotel *types.Hotel
	if err := s.coll.FindOne(ctx, filter).Decode(&hotel); err != nil {
		return nil, err
	}
	
	return hotel, nil
}