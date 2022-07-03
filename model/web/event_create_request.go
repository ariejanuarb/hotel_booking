package web

type EventCreateRequest struct {
	Event_start string `validate:"required" json:"event_Start"`
	Event_end   string `validate:"required" json:"event_End"`
}
