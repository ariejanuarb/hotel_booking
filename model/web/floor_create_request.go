package web

type FloorCreateRequest struct {
	Number   int `validate:"required" json:"number"`
	Capacity int `validate:"required" json:"capacity"`
}
