package repository

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
)

type BookingRepository interface {
	Save(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking
	Update(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking
	FindById(ctx context.Context, tx *sql.Tx, bookingId int) (web.BookingResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []web.BookingResponse
}
