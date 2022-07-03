package web

import "time"

type EventResponse struct {
	Event_id    int       `json:"event_Id"`
	Event_start string    `json:"event_Start"`
	Event_end   string    `json:"event_End"`
	Created_at  time.Time `json:"created_At"`
	Updated_at  time.Time `json:"updated_At"`
}
