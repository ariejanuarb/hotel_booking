package web

type RolesUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=100,min=1" json:"number"`
}
