package service

import (
	"context"
	"hotel_booking/model/web"
)

type FloorService interface {
	Create(ctx context.Context, request web.FloorCreateRequest) web.FloorResponse
	Update(ctx context.Context, request web.FloorUpdateRequest) web.FloorResponse
	Delete(ctx context.Context, floorId int)
	FindById(ctx context.Context, floorId int) web.FloorResponse
	FindAll(ctx context.Context) []web.FloorResponse
}
