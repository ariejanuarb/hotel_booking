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

func (controller *InvoiceControllerImpl) CreateInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceCreateRequest := web.InvoiceCreateRequest{}
	helper.ReadFromRequestBody(request, &invoiceCreateRequest)

	invoiceResponse := controller.InvoiceService.CreateInvoice(request.Context(), invoiceCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) UpdateInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceUpdateRequest := web.InvoiceUpdateRequest{}
	helper.ReadFromRequestBody(request, &invoiceUpdateRequest)

	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	invoiceUpdateRequest.Invoice_id = id

	invoiceResponse := controller.InvoiceService.UpdateInvoice(request.Context(), invoiceUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) DeleteInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	controller.InvoiceService.DeleteInvoice(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) FindInvoiceById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	invoiceResponse := controller.InvoiceService.FindInvoiceById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) FindAllInvoice(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceResponse := controller.InvoiceService.FindAllInvoice(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
