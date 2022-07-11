package helper

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
)

func ToBookingResponse(b domain.Booking) web.BookingResponse {
	return web.BookingResponse{
		Id:                 b.Id,
		Status:             b.Status,
		Room_id:            b.Room_id,
		Pic_name:           b.Pic_name,
		Pic_Contact:        b.Pic_Contact,
		Event_start:        b.Event_start,
		Event_end:          b.Event_end,
		Invoice_number:     b.Invoice_number,
		Invoice_subtotal:   b.Invoice_subtotal,
		Invoice_grandtotal: b.Invoice_grandtotal,
		Discount_amount:    b.Discount_amount,
		Discount_request:   b.Discount_request,
		Created_at:         b.Created_at,
		Updated_at:         b.Updated_at,
	}
}

func ToBooking(b web.BookingResponse) domain.Booking {
	return domain.Booking{
		Id:                 b.Id,
		Status:             b.Status,
		Room_id:            b.Room_id,
		Pic_name:           b.Pic_name,
		Pic_Contact:        b.Pic_Contact,
		Event_start:        b.Event_start,
		Event_end:          b.Event_end,
		Invoice_number:     b.Invoice_number,
		Invoice_subtotal:   b.Invoice_subtotal,
		Invoice_grandtotal: b.Invoice_grandtotal,
		Discount_amount:    b.Discount_amount,
		Discount_request:   b.Discount_request,
		Created_at:         b.Created_at,
		Updated_at:         b.Updated_at,
	}
}

func ToBookingResponses(b []domain.Booking) []web.BookingResponse {
	var bookingResponses []web.BookingResponse
	for _, booking := range b {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}
	return bookingResponses
}
