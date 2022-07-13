package service

import (
	"booking-hotel/model/web"
	"context"
)

type BookingService interface {
	Create(ctx context.Context, request *web.BookingCreateRequest) (*web.BookingResponse, error)
	UpdateStatus(ctx context.Context, request *web.UpdateStatus) *web.BookingResponse
	UpdateDiscount(ctx context.Context, request *web.UpdateDiscount) *web.BookingResponse
	FindById(ctx context.Context, bookingId int) *web.BookingResponse
	FindAll(ctx context.Context) []*web.BookingResponse
}
