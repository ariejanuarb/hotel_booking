package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type DiscountControllerImpl struct {
	DiscountService service.DiscountService
}

func NewDiscountController(discountService service.DiscountService) DiscountController {
	return &DiscountControllerImpl{
		DiscountService: discountService,
	}
}

func (controller *DiscountControllerImpl) CreateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountCreateRequest := web.DiscountCreateRequest{}
	helper.ReadFromRequestBody(request, &discountCreateRequest)

	discountResponse := controller.DiscountService.CreateDiscount(request.Context(), discountCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) UpdateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountUpdateRequest := web.DiscountUpdateRequest{}
	helper.ReadFromRequestBody(request, &discountUpdateRequest)

	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	discountUpdateRequest.Discount_id = id

	discountResponse := controller.DiscountService.UpdateDiscount(request.Context(), discountUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) DeleteDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	controller.DiscountService.DeleteDiscount(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) FindDiscountById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	discountResponse := controller.DiscountService.FindDiscountById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DiscountControllerImpl) FindAllDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountResponse := controller.DiscountService.FindAllDiscount(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
