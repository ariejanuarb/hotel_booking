package web

import "time"

type BookingResponse struct {
	Id                 int       `json:"id"`
	Status             string    `json:"status"`
	Room_id            int       `json:"room_Id"`
	Room_name          string    `json:"room_Name"`
	Pic_name           string    `json:"pic_Name"`
	Pic_Contact        string    `json:"pic_Contact"`
	Event_start        string    `json:"event_Start"`
	Event_end          string    `json:"event_End"`
	Invoice_number     string    `json:"invoice_Number"`
	Invoice_subtotal   string    `json:"invoice_Subtotal"`
	Invoice_grandtotal string    `json:"invoice_Grandtotal"`
	Discount_request   string    `json:"discount_Request"`
	Discount_amount    string    `json:"discount_Amount"`
	Created_at         time.Time `json:"created_At"`
	Updated_at         time.Time `json:"updated_At"`
}
