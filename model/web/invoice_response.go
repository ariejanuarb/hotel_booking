package web

import "time"

type InvoiceResponse struct {
	Invoice_id     int       `json:"invoice_Id"`
	Invoice_Date   string    `json:"invoice_Date"`
	Tax            string    `json:"tax"`
	Price          string    `json:"price"`
	Invoice_number int       `json:"invoice_Number"`
	Total          string    `json:"total"`
	Created_at     time.Time `json:"created_At"`
	Updated_at     time.Time `json:"updated_At"`
}
