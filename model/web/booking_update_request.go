package web

type BookingUpdateRequest struct {
	Id     int    `validate:"required" json:"id"`
	Status string `validate:"required,min=1,max=100" json:"status"`
}
