package domain

import "time"

type Discount struct {
	Discount_id      int
	Discount_request string
	Discount_status  string
	Discount_amount  float64
	Invoice_id       int
	Created_at       time.Time
	Updated_at       time.Time
}
