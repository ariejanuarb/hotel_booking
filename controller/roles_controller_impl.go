package controller

import (
	"github.com/julienschmidt/httprouter"
	"hotel_booking/helper"
	"hotel_booking/model/web"
	"hotel_booking/service"
	"net/http"
	"strconv"
)

type RolesControllerImpl struct {
	RolesService service.RolesService
}

func NewRolesController(rolesService service.RolesService) RolesController {
	return &RolesControllerImpl{
		RolesService: rolesService,
	}
}

func (controller *RolesControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesCreateRequest := web.RolesCreateRequest{}
	helper.ReadFromRequestBody(request, &rolesCreateRequest)

	rolesResponse := controller.RolesService.Create(request.Context(), rolesCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rolesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RolesControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesUpdateRequest := web.RolesUpdateRequest{}
	helper.ReadFromRequestBody(request, &rolesUpdateRequest)

	rolesId := params.ByName("rolesId")
	id, err := strconv.Atoi(rolesId)
	helper.PanicIfError(err)

	rolesUpdateRequest.Id = id

	rolesResponse := controller.RolesService.Update(request.Context(), rolesUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rolesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RolesControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesId := params.ByName("rolesId")
	id, err := strconv.Atoi(rolesId)
	helper.PanicIfError(err)

	controller.RolesService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RolesControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesId := params.ByName("rolesId")
	id, err := strconv.Atoi(rolesId)
	helper.PanicIfError(err)

	rolesResponse := controller.RolesService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rolesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RolesControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesResponses := controller.RolesService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rolesResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
