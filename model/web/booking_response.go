package web

import "time"

type BookingResponse struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Pic_name    string    `json:"pic_Name"`
	Pic_Contact string    `json:"pic_Contact"`
	Invoice     string    `json:"invoice"`
	Event_start string    `json:"event_Start"`
	Event_end   string    `json:"event_End"`
	Room_id     int       `json:"room_Id"`
	Room_name   string    `json:"room_Name"`
	Created_at  time.Time `json:"created_At"`
	Updated_at  time.Time `json:"updated_At"`
}
