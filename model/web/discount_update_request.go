package web

type DiscountUpdateRequest struct {
	Discount_id      int     `validate:"required"`
	Discount_request string  `validate:"required,min=1,max=100" json:"discount_request"`
	Discount_status  string  `validate:"required,min=1,max=100" json:"discount_Status"`
	Discount_amount  float64 `validate:"required,min=1" json:"discount_Amount"`
	Invoice_id       int     `validate:"required,min=1,max=100" json:"invoice_Id"`
}
