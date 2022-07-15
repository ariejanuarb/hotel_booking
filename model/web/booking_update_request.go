package web

type UpdateRequest struct {
	Id     int    `validate:"required" json:"id"`
	Status string `json:"status"`
}

type UpdateDiscount struct {
	Id               int    `validate:"required" json:"id"`
	Discount_request string `json:"discount_Request"`
	Discount_amount  string `json:"discount_Amount"`
}

type ResponseDiscount struct {
	Id               int    `validate:"required" json:"id"`
	Discount_request string `validate:"required,min=1" json:"discount_Request"`
	Discount_amount  string `validate:"required,min=1" json:"discount_Amount"`
}
