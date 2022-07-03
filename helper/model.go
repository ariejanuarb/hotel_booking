package helper

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"time"
)

func ToDiscountResponse(disc domain.Discount) web.DiscountResponse {
	return web.DiscountResponse{
		Discount_id:      disc.Discount_id,
		Discount_request: disc.Discount_request,
		Discount_status:  disc.Discount_status,
		Discount_amount:  disc.Discount_amount,
		Invoice_id:       disc.Invoice_id,
		Created_at:       time.Now(),
		Updated_at:       time.Now(),
	}
}

func ToDiscount(disc web.DiscountResponse) domain.Discount {
	return domain.Discount{
		Discount_id:      disc.Discount_id,
		Discount_request: disc.Discount_request,
		Discount_status:  disc.Discount_status,
		Discount_amount:  disc.Discount_amount,
		Invoice_id:       disc.Invoice_id,
		Created_at:       time.Now(),
		Updated_at:       time.Now(),
	}
}

func ToDiscountResponses(disc []domain.Discount) []web.DiscountResponse {
	var discountResponses []web.DiscountResponse
	for _, discount := range disc {
		discountResponses = append(discountResponses, ToDiscountResponse(discount))
	}
	return discountResponses
}

func ToEventResponse(event domain.Event) web.EventResponse {
	return web.EventResponse{
		Event_id:    event.Event_id,
		Event_start: event.Event_start,
		Event_end:   event.Event_End,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
}

func ToEventResponses(event []domain.Event) []web.EventResponse {
	var eventResponses []web.EventResponse
	for _, events := range event {
		eventResponses = append(eventResponses, ToEventResponse(events))
	}
	return eventResponses
}

func ToInvoiceResponse(invcs domain.Invoice) web.InvoiceResponse {
	return web.InvoiceResponse{
		Invoice_id:   invcs.Invoice_id,
		Invoice_Date: invcs.Invoice_Date,
		Tax:          invcs.Tax,
		Price:        invcs.Price,
		Total:        invcs.Total,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}
}

func ToInvoiceResponses(invcs []domain.Invoice) []web.InvoiceResponse {
	var invoiceResponses []web.InvoiceResponse
	for _, invoices := range invcs {
		invoiceResponses = append(invoiceResponses, ToInvoiceResponse(invoices))
	}
	return invoiceResponses
}
