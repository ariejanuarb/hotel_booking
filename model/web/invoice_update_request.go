package web

type InvoiceUpdateRequest struct {
	Invoice_id   int    `validate:"required"`
	Invoice_date string `validate:"required,min=1" json:"invoice_Date"`
	Tax          string `validate:"required,min=1" json:"tax"`
	Price        string `validate:"required,min=1" json:"price"`
	Total        string `validate:"required,min=1" json:"total"`
}
