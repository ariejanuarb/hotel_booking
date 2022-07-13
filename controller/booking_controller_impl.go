package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookingControllerImpl struct {
	BookingService service.BookingService
}

func NewBookingController(BookingService *service.BookingServiceImpl) BookingController {
	return &BookingControllerImpl{
		BookingService: BookingService,
	}
}

func (controller *BookingControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var webResponse web.WebResponse
	bookingCreateRequest := web.BookingCreateRequest{}
	helper.ReadFromRequestBody(request, &bookingCreateRequest)

	bookingResponse, err := controller.BookingService.Create(request.Context(), &bookingCreateRequest)
	if err != nil {
		webResponse = web.WebResponse{
			Code:   400,
			Status: "Error",
			Data:   err.Error(),
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "Ok",
			Data:   bookingResponse,
		}
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateStatusRequest := web.UpdateRequest{}
	helper.ReadFromRequestBody(request, &updateStatusRequest)

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	updateStatusRequest.Id = id

	bookingResponse := controller.BookingService.UpdateStatus(request.Context(), &updateStatusRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) UpdateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateDiscountRequest := web.UpdateRequest{}
	helper.ReadFromRequestBody(request, &updateDiscountRequest)

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	updateDiscountRequest.Id = id

	bookingResponse := controller.BookingService.UpdateDiscount(request.Context(), &updateDiscountRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) ResponseDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responseDiscountRequest := web.UpdateRequest{}
	helper.ReadFromRequestBody(request, &responseDiscountRequest)

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	responseDiscountRequest.Id = id

	bookingResponse := controller.BookingService.UpdateDiscount(request.Context(), &responseDiscountRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindAllDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingResponse := controller.BookingService.FindAllDiscount(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	bookingResponse := controller.BookingService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingResponse := controller.BookingService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
