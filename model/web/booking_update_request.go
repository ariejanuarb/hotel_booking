package web

type UpdateStatus struct {
	Id     int    `validate:"required" json:"id"`
	Status string `validate:"required,min=1,max=100" json:"status"`
}

type UpdateDiscount struct {
	Id               int    `validate:"required" json:"id"`
	Discount_request string `validate:"required,min=1" json:"discount_Request"`
	Discount_amount  string `validate:"required,min=1" json:"discount_Amount"`
}
