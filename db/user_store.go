package db

import (
	"context"
	"fmt"
	"golang-hotel-reservation/types"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// the bson.M is mongodb specific, GeneralizedBson is a more generalized for other db providers
type GeneralizedBson map[string]any

const userColl = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error )
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter GeneralizedBson, update GeneralizedBson) error
	GetUserByEmail(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(userColl),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("... dropping user collection")
	return s.coll.Drop(ctx)
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err !=nil {
		return nil, err
	}

	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return users, err
	}

	return users, nil
}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
} 

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// TODO: maybe it's a good idea to handle if we did not delete any user.
	// maybe log it or something??
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter GeneralizedBson, update GeneralizedBson) error {
	oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
	if err != nil {
		return err
	}
	filter["_id"] = oid
	_, err = s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}