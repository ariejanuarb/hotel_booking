package service

import (
	"context"
	"hotel_booking/model/web"
)

type RolesService interface {
	Create(ctx context.Context, request web.RolesCreateRequest) web.RolesResponse
	Update(ctx context.Context, request web.RolesUpdateRequest) web.RolesResponse
	Delete(ctx context.Context, rolesId int)
	FindById(ctx context.Context, rolesId int) web.RolesResponse
	FindAll(ctx context.Context) []web.RolesResponse
}
