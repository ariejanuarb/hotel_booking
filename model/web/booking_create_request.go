package web

type BookingCreateRequest struct {
	Status      string `validate:"required,min=1,max=100" json:"status"`
	Room_id     int    `validate:"required,min=1" json:"room_Id"`
	Pic_name    string `validate:"required,min=1,max=100" json:"pic_Name"`
	Pic_Contact string `validate:"required,min=1,max=100" json:"pic_Contact"`
	Event_start string `validate:"required,min=1,max=100" json:"event_Start"`
	Event_end   string `validate:"required,min=1,max=100" json:"event_End"`
}
