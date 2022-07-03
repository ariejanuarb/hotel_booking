package domain

import "time"

type Event struct {
	Event_id    int
	Event_start string
	Event_End   string
	Created_at  time.Time
	Updated_at  time.Time
}
