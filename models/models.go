package models

import (
	"applicationDesignTest/utils"
	"time"
)

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

var Orders = []Order{}

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

var Availability = []RoomAvailability{
	{HotelID: "reddison", RoomID: "lux", Date: utils.Date(2024, 1, 1), Quota: 1},
	{"reddison", "lux", utils.Date(2024, 1, 2), 1},
	{"reddison", "lux", utils.Date(2024, 1, 3), 1},
	{"reddison", "lux", utils.Date(2024, 1, 4), 1},
	{"reddison", "lux", utils.Date(2024, 1, 5), 0},
}
