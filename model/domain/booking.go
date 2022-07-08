package domain

import "time"

type Booking struct {
	Id          int
	Status      string
	Pic_name    string
	Pic_Contact string
	Invoice     string
	Event_start string
	Event_end   string
	Room_id     int
	Created_at  time.Time
	Updated_at  time.Time
}
