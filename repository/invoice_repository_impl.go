package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type InvoiceRepositoryImpl struct {
}

func NewInvoiceRepository() InvoiceRepository {
	return &InvoiceRepositoryImpl{}
}

func (i InvoiceRepositoryImpl) SaveInvoice(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice {
	SQL := "insert into invoice(invoice_date, tax, price, total) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, invoice.Invoice_Date, invoice.Tax, invoice.Price, invoice.Total)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	invoice.Invoice_id = int(id)
	return invoice
}

func (i InvoiceRepositoryImpl) UpdateInvoice(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice {
	SQl := "update invoice set invoice_date = ?, tax = ?, price = ?, total = ? where invoice_id = ?"
	_, err := tx.ExecContext(ctx, SQl, invoice.Invoice_Date, invoice.Tax, invoice.Price, invoice.Total, invoice.Invoice_id)
	helper.PanicIfError(err)

	return invoice
}

func (i InvoiceRepositoryImpl) DeleteInvoice(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) {
	SQL := "delete from invoice where invoice_id = ?"
	_, err := tx.ExecContext(ctx, SQL, invoice.Invoice_id)
	helper.PanicIfError(err)
}

func (i InvoiceRepositoryImpl) FindInvoiceById(ctx context.Context, tx *sql.Tx, invoiceId int) (domain.Invoice, error) {
	SQL := "select invoice_id, invoice_date, tax, price, total from invoice where invoice_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, invoiceId)
	helper.PanicIfError(err)
	defer rows.Close()

	invoices := domain.Invoice{}

	if rows.Next() {
		err := rows.Scan(&invoices.Invoice_id, &invoices.Invoice_Date, &invoices.Tax, &invoices.Price, &invoices.Total)
		helper.PanicIfError(err)
		defer rows.Close()
		return invoices, nil
	} else {
		return invoices, errors.New("invoice is not found")
	}

}

func (i InvoiceRepositoryImpl) FindAllInvoice(ctx context.Context, tx *sql.Tx) []domain.Invoice {
	SQL := "select invoice_id, invoice_date, tax, price, total from invoice"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var invoice []domain.Invoice
	for rows.Next() {
		invoices := domain.Invoice{}
		err := rows.Scan(&invoices.Invoice_id, &invoices.Invoice_Date, &invoices.Tax, &invoices.Price, &invoices.Total)
		helper.PanicIfError(err)
		invoice = append(invoice, invoices)
	}
	return invoice
}
