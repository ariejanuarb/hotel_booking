package service

import (
	"booking-hotel/model/web"
	"context"
)

type BookingService interface {
	Create(ctx context.Context, request web.BookingCreateRequest) web.BookingResponse
	Update(ctx context.Context, request web.BookingUpdateRequest) web.BookingResponse
	FindById(ctx context.Context, bookingId int) web.BookingResponse
	FindAll(ctx context.Context) []web.BookingResponse
}
