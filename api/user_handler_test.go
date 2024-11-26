package api

import (
	"bytes"
	"encoding/json"
	"golang-hotel-reservation/types"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)





func TestPostUser(t *testing.T)  {
	tdb := setup(t)
	defer tdb.teardown(t)
	
	app := fiber.New()
	UserHandler := NewUserHandler(tdb.User)
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