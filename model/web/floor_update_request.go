package web

type FloorUpdateRequest struct {
	Id       int `validate:"required"`
	Number   int `validate:"required" json:"number"`
	Capacity int `validate:"required" json:"capacity"`
}
