package domain

import "time"

type Booking struct {
	Id                 int
	Status             string
	Room_id            int
	Pic_name           string
	Pic_Contact        string
	Event_start        string
	Event_end          string
	Invoice_number     int
	Invoice_subtotal   string
	Invoice_grandtotal string
	Discount_request   string
	Discount_amount    string
	Created_at         time.Time
	Updated_at         time.Time
}
