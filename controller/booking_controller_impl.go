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
	bookingCreateRequest := web.BookingCreateRequest{}
	helper.ReadFromRequestBody(request, &bookingCreateRequest)

	bookingResponse, _ := controller.BookingService.Create(request.Context(), &bookingCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UpdateStatusRequest := web.UpdateStatus{}
	helper.ReadFromRequestBody(request, &UpdateStatusRequest)

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	UpdateStatusRequest.Id = id

	bookingResponse := controller.BookingService.UpdateStatus(request.Context(), &UpdateStatusRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) UpdateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateDiscountRequest := web.UpdateDiscount{}
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
