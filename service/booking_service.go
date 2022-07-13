package service

import (
	"booking-hotel/model/web"
	"context"
)

type BookingService interface {
	Create(ctx context.Context, request *web.BookingCreateRequest) (*web.BookingResponse, error)
	UpdateStatus(ctx context.Context, request *web.UpdateRequest) *web.BookingResponse
	UpdateDiscount(ctx context.Context, request *web.UpdateRequest) *web.BookingResponse
	ResponseDiscount(ctx context.Context, request *web.UpdateRequest) *web.BookingResponse
	FindAllDiscount(ctx context.Context) []*web.BookingResponse
	FindById(ctx context.Context, bookingId int) *web.BookingResponse
	FindAll(ctx context.Context) []*web.BookingResponse
}
