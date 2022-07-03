package web

type EventUpdateRequest struct {
	Event_id    int    `validate:"required,min=1" json:"event_Id"`
	Event_start string `validate:"required" json:"event_Start"`
	Event_end   string `validate:"required" json:"event_End"`
}
