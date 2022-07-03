package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type InvoiceRepository interface {
	SaveInvoice(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice
	UpdateInvoice(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice
	DeleteInvoice(ctx context.Context, tx *sql.Tx, invoice domain.Invoice)
	FindInvoiceById(ctx context.Context, tx *sql.Tx, invoiceId int) (domain.Invoice, error)
	FindAllInvoice(ctx context.Context, tx *sql.Tx) []domain.Invoice
}
