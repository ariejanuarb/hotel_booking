package helper

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"time"
)

func ToBookingResponse(b domain.Booking) web.BookingResponse {
	return web.BookingResponse{
		Id:          b.Id,
		Status:      b.Status,
		Pic_name:    b.Pic_name,
		Pic_Contact: b.Pic_Contact,
		Invoice:     b.Invoice,
		Event_start: b.Event_start,
		Event_end:   b.Event_end,
		Room_id:     b.Room_id,
		Created_at:  time.Now(),
	}
}

func ToBooking(b web.BookingResponse) domain.Booking {
	return domain.Booking{
		Id:          b.Id,
		Status:      b.Status,
		Pic_name:    b.Pic_name,
		Pic_Contact: b.Pic_Contact,
		Invoice:     b.Invoice,
		Event_start: b.Event_start,
		Event_end:   b.Event_end,
		Room_id:     b.Room_id,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
}

func ToBookingResponses(b []domain.Booking) []web.BookingResponse {
	var bookingResponses []web.BookingResponse
	for _, booking := range b {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}
	return bookingResponses
}
