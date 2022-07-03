package web

type DiscountCreateRequest struct {
	Discount_request string  `validate:"required,min=1,max=100" json:"discount_Request"`
	Discount_status  string  `validate:"required,min=1,max=100" json:"discount_Status"`
	Discount_amount  float64 `valudate:"min=1" json:"discount_Amount"`
	Invoice_id       int     `validate:"required,min=1,max=100" json:"invoice_Id"`
}
