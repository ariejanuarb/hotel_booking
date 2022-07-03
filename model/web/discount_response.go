package web

import "time"

type DiscountResponse struct {
	Discount_id      int       `json:"discount_id"`
	Discount_request string    `json:"discount_request"`
	Discount_status  string    `json:"discount_Status"`
	Discount_amount  float64   `json:"discount_Amount"`
	Invoice_id       int       `json:"invoice_Id"`
	Created_at       time.Time `json:"created_At"`
	Updated_at       time.Time `json:"updated_At"`
}
