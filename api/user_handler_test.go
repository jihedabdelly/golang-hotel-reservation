package api

import (
	"bytes"
	"context"
	"encoding/json"
	"golang-hotel-reservation/db"
	"golang-hotel-reservation/types"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
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
	
	app := fiber.New()
	UserHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", UserHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email: "some@email.com",
		FirstName: "cristiano",
		LastName: "rolando",
		Password: "sfgzgthzghtzsdgfhb",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0  {
		t.Errorf("Expecting a user id to be set")
	}
	if len(user.EncryptedPassword) > 0  {
		t.Errorf("Expected the user password not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected the firstname to be %s but got %s", params.FirstName, user.FirstName)
	}
  if user.LastName != params.LastName {
		t.Errorf("Expected the lastname to be %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("Expected the email to be %s but got %s", params.Email, user.Email)
	}
	
}