package api

import (
	"golang-hotel-reservation/db"
	"golang-hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hs,
		roomStore:  rs,
	}
}

type HotelQueryParams struct {
	Rooms bool
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var queryParams HotelQueryParams
	if err := c.QueryParser(&queryParams); err != nil {
		return err
	}

	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	if  queryParams.Rating > 0 {
		return c.JSON(hotels)
	}

	var filteredHotels []types.Hotel
	for _, hotel := range hotels {
		if hotel.Rating == queryParams.Rating {
			filteredHotels = append(filteredHotels, *hotel)
		}
	}

	return c.JSON(filteredHotels)
}
