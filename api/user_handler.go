package api

import (
	"errors"
	"golang-hotel-reservation/db"
	"golang-hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found!"})
		}
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), *user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	err := h.userStore.DeleteUser(c.Context(), userId)
	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"deleted": userId})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		userId = c.Params("id")
		values types.UpdateUserParams
	)

	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}

	if err := c.BodyParser(&values); err != nil {
		return err
	}
	update := bson.M{"$set": values}

	if err = h.userStore.UpdateUser(c.Context(), filter, update); err != nil {
		return err
	}

	return c.JSON(map[string]string{"updated": userId})
}
