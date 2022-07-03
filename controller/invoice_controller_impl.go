package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type InvoiceControllerImpl struct {
	InvoiceService service.InvoiceService
}

func NewInvoiceController(invoiceService service.InvoiceService) InvoiceController {
	return &InvoiceControllerImpl{
		InvoiceService: invoiceService,
	}
}

func (i InvoiceControllerImpl) CreateInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceCreateRequest := web.InvoiceCreateRequest{}
	helper.ReadFromRequestBody(request, &invoiceCreateRequest)

	invoiceResponse := i.InvoiceService.CreateInvoice(request.Context(), invoiceCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (i InvoiceControllerImpl) UpdateInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceUpdateRequest := web.InvoiceUpdateRequest{}
	helper.ReadFromRequestBody(request, &invoiceUpdateRequest)

	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	invoiceUpdateRequest.Invoice_id = id

	invoiceResponse := i.InvoiceService.UpdateInvoice(request.Context(), invoiceUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (i InvoiceControllerImpl) DeleteInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	i.InvoiceService.DeleteInvoice(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (i InvoiceControllerImpl) FindInvoiceById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	invoiceResponse := i.InvoiceService.FindInvoiceById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (i InvoiceControllerImpl) FindAllInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceResponse := i.InvoiceService.FindAllInvoice(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
