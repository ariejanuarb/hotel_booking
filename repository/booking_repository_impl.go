package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
	"errors"
)

type BookingRepositoryImpl struct {
}

func NewBookingRepository() BookingRepository {
	return &BookingRepositoryImpl{}
}

func (repository *BookingRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking {
	SQL := "insert into booking(status, pic_name, pic_contact, invoice, event_start, event_end, room_id) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, booking.Status, booking.Pic_name, booking.Pic_Contact, booking.Invoice, booking.Event_start, booking.Event_end, booking.Room_id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	booking.Id = int(id)
	return booking
}

func (repository *BookingRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, booking domain.Booking) domain.Booking {
	SQL := "update booking set status = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Status, booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository *BookingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookingId int) (web.BookingResponse, error) {
	SQL := "select b.id, b.status, b.pic_name, b.pic_contact, b.invoice, b.event_start, b.event_end, b.room_id, r.name as 'room_name' from booking b inner join room r on b.room_id=r.id where b.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookingId)
	helper.PanicIfError(err)
	defer rows.Close()

	booking := web.BookingResponse{}
	if rows.Next() {
		err := rows.Scan(&booking.Id, &booking.Status, &booking.Pic_name, &booking.Pic_Contact, &booking.Invoice, &booking.Event_start, &booking.Event_end, &booking.Room_id, &booking.Room_name)
		helper.PanicIfError(err)
		defer rows.Close()
		return booking, nil
	} else {
		return booking, errors.New("Booking is not found")
	}
}

func (repository *BookingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.BookingResponse {
	SQL := "select b.id, b.status, b.pic_name, b.pic_contact, b.invoice, b.event_start, b.event_end, b.room_id, r.name as 'room_name' from booking b inner join room r on b.room_id=r.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var bookings []web.BookingResponse
	for rows.Next() {
		booking := web.BookingResponse{}
		err := rows.Scan(&booking.Id, &booking.Status, &booking.Pic_name, &booking.Pic_Contact, &booking.Invoice, &booking.Event_start, &booking.Event_end, &booking.Room_id, &booking.Room_name)
		helper.PanicIfError(err)
		bookings = append(bookings, booking)
	}
	return bookings
}
