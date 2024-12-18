package api

import (
	"golang-hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter :=  db.GeneralizedBson{
		"rating": params.Rating,
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrResourceNotFound("hotels")
	}
	resp := ResourceResp{
		Data: hotels,
		Results: len(hotels),
		Page: int(params.Page),
	}
	return c.JSON(resp)
}



func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := db.GeneralizedBson{"_id": oid}
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), filter)
	if err != nil {
		return ErrResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := db.GeneralizedBson{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrResourceNotFound("rooms")
	}

	return c.JSON(rooms)
}
