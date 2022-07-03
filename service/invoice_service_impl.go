package service

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator"
)

type InvoiceServiceImpl struct {
	InvoiceRepository repository.InvoiceRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository, db *sql.DB, validate *validator.Validate) InvoiceService {
	return &InvoiceServiceImpl{
		InvoiceRepository: invoiceRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (i InvoiceServiceImpl) CreateInvoice(ctx context.Context, request web.InvoiceCreateRequest) web.InvoiceResponse {
	err := i.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := i.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	invoice := domain.Invoice{
		Invoice_Date: request.Invoice_date,
		Tax:          request.Tax,
		Price:        request.Price,
		Total:        request.Total,
	}
	invoice = i.InvoiceRepository.SaveInvoice(ctx, tx, invoice)
	return helper.ToInvoiceResponse(invoice)
}

func (i InvoiceServiceImpl) UpdateInvoice(ctx context.Context, request web.InvoiceUpdateRequest) web.InvoiceResponse {
	err := i.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := i.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	invoice, err := i.InvoiceRepository.FindInvoiceById(ctx, tx, request.Invoice_id)
	if err != nil {
		helper.PanicIfError(err)
	}
	invoice.Invoice_Date = request.Invoice_date
	invoice.Tax = request.Tax
	invoice.Price = request.Price
	invoice.Total = request.Total

	invoice = i.InvoiceRepository.UpdateInvoice(ctx, tx, invoice)
	return helper.ToInvoiceResponse(invoice)
}

func (i InvoiceServiceImpl) DeleteInvoice(ctx context.Context, invoiceId int) {
	tx, err := i.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	invoice, err := i.InvoiceRepository.FindInvoiceById(ctx, tx, invoiceId)
	if err != nil {
		helper.PanicIfError(err)
	}
	i.InvoiceRepository.DeleteInvoice(ctx, tx, invoice)
}

func (i InvoiceServiceImpl) FindInvoiceById(ctx context.Context, invoiceId int) web.InvoiceResponse {
	tx, err := i.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	invoice, err := i.InvoiceRepository.FindInvoiceById(ctx, tx, invoiceId)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper.ToInvoiceResponse(invoice)
}

func (i InvoiceServiceImpl) FindAllInvoice(ctx context.Context) []web.InvoiceResponse {
	tx, err := i.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	invoices := i.InvoiceRepository.FindAllInvoice(ctx, tx)
	return helper.ToInvoiceResponses(invoices)
}
