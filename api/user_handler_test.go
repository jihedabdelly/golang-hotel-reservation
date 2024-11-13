package api

import (
	"context"
	"golang-hotel-reservation/db"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type testdb struct {
	db.UserStore
}

const testdburi = "mongodb://localhost:27017"

func (tbd *testdb) teardown(t *testing.T) {
	if err := tbd.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, db.DBNAME_TEST),
	}

}

func TestPostUser(t *testing.T)  {
	tdb := setup(t)
	defer tdb.teardown(t)
	//t.Fail()
}