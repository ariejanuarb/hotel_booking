package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type InvoiceController interface {
	CreateInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindInvoiceById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
