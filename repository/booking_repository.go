package repository

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
)

type BookingRepository interface {
	Save(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	UpdateStatus(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	UpdateDiscount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	ResponseDiscount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	FindAllDiscount(ctx context.Context, tx *sql.Tx) []*web.BookingResponse
	FindById(ctx context.Context, tx *sql.Tx, bookingId int) (*web.BookingResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []*web.BookingResponse
}
