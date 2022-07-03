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

func (disc DiscountControllerImpl) CreateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountCreateRequest := web.DiscountCreateRequest{}
	helper.ReadFromRequestBody(request, &discountCreateRequest)

	discountResponse := disc.DiscountService.CreateDiscount(request.Context(), discountCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (disc DiscountControllerImpl) UpdateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountUpdateRequest := web.DiscountUpdateRequest{}
	helper.ReadFromRequestBody(request, &discountUpdateRequest)

	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	discountUpdateRequest.Discount_id = id

	discountResponse := disc.DiscountService.UpdateDiscount(request.Context(), discountUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (disc DiscountControllerImpl) DeleteDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	disc.DiscountService.DeleteDiscount(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (disc DiscountControllerImpl) FindDiscountById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountId := params.ByName("discountId")
	id, err := strconv.Atoi(discountId)
	helper.PanicIfError(err)

	discountResponse := disc.DiscountService.FindDiscountById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (disc DiscountControllerImpl) FindAllDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	discountResponse := disc.DiscountService.FindAllDiscount(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   discountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
