package web

type UpdateRequest struct {
	Id               int    `validate:"required" json:"id"`
	Status           string `validate:"required" json:"status"`
	Discount_request string `validate:"required,min=1" json:"discount_Request"`
	Discount_amount  string `validate:"required,min=1" json:"discount_Amount"`
}
