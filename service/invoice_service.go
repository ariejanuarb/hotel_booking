package service

import (
	"booking-hotel/model/web"
	"context"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, request web.InvoiceCreateRequest) web.InvoiceResponse
	UpdateInvoice(ctx context.Context, request web.InvoiceUpdateRequest) web.InvoiceResponse
	DeleteInvoice(ctx context.Context, invoiceId int)
	FindInvoiceById(ctx context.Context, invoiceId int) web.InvoiceResponse
	FindAllInvoice(ctx context.Context) []web.InvoiceResponse
}
