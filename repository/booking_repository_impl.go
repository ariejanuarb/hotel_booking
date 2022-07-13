package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
	"errors"
	"time"
)

type BookingRepositoryImpl struct {
}

func NewBookingRepository() BookingRepository {
	return &BookingRepositoryImpl{}
}

func (repository *BookingRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "insert into booking(status, room_id, pic_name, pic_contact, event_start, event_end, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, booking.Status, booking.Room_id, booking.Pic_name, booking.Pic_Contact, booking.Event_start, booking.Event_end, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	booking.Id = int(id)
	return booking
}

func (repository *BookingRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "update booking set status = ?, updated_at= ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Status, time.Now().Format("2006-01-02 15:04:05"), booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository *BookingRepositoryImpl) UpdateDiscount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "update booking set discount_request = ?, discount_amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Discount_request, booking.Discount_amount, booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository *BookingRepositoryImpl) GetEvent(ctx context.Context, tx *sql.Tx, bookingId int) (*web.BookingResponse, error) {
	SQL := "select event_start, event_end, room_id from booking where id = *"
	rows, err := tx.QueryContext(ctx, SQL, bookingId)
	helper.PanicIfError(err)
	defer rows.Close()

	booking := web.BookingResponse{}
	if rows.Next() {
		err := rows.Scan(&booking.Event_start, &booking.Event_end, &booking.Room_id)
		helper.PanicIfError(err)
		defer rows.Close()
		return &booking, nil
	} else {
		return &booking, errors.New("booking is not found")
	}
}

func (repository *BookingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookingId int) (*web.BookingResponse, error) {
	SQL := "select b.id, b.status, b.room_id, b.pic_name, b.pic_contact, b.event_start, b.event_end, b.invoice_number, b.invoice_subtotal, b.invoice_grandtotal, b.discount_request, b.discount_amount, r.price_per_hour, r.price_per_day from booking b inner join room r on b.room_id=r.id where b.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookingId)
	helper.PanicIfError(err)
	defer rows.Close()

	booking := web.BookingResponse{}
	if rows.Next() {
		err := rows.Scan(&booking.Id, &booking.Status, &booking.Room_id, &booking.Pic_name, &booking.Pic_Contact, &booking.Event_start, &booking.Event_end, &booking.Invoice_number, &booking.Invoice_subtotal, &booking.Invoice_grandtotal, &booking.Discount_request, &booking.Discount_amount, &booking.Price_per_hour, &booking.Price_per_day)
		helper.PanicIfError(err)
		defer rows.Close()
		return &booking, nil
	} else {
		return &booking, errors.New("Booking is not found")
	}
}

func (repository *BookingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*web.BookingResponse {
	SQL := "select b.id, b.status, b.room_id, b.pic_name, b.pic_contact, b.event_start, b.event_end, b.invoice_number, b.invoice_subtotal, b.invoice_grandtotal, b.discount_request, b.discount_amount, r.price_per_hour, r.price_per_day from booking b inner join room r on b.room_id=r.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var bookings []*web.BookingResponse
	for rows.Next() {
		booking := &web.BookingResponse{}
		err := rows.Scan(&booking.Id, &booking.Status, &booking.Room_id, &booking.Pic_name, &booking.Pic_Contact, &booking.Event_start, &booking.Event_end, &booking.Invoice_number, &booking.Invoice_subtotal, &booking.Invoice_grandtotal, &booking.Discount_request, &booking.Discount_amount, &booking.Price_per_hour, &booking.Price_per_day)
		helper.PanicIfError(err)
		bookings = append(bookings, booking)
	}
	return bookings
}
