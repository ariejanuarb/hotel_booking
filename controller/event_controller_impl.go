package controller

import (
	"booking-hotel/helper"
	"booking-hotel/model/web"
	"booking-hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func NewEventController(eventService service.EventService) EventController {
	return &EventControllerImpl{
		EventService: eventService,
	}
}

func (e EventControllerImpl) CreateEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventCreateRequest := web.EventCreateRequest{}
	helper.ReadFromRequestBody(request, &eventCreateRequest)

	eventResponse := e.EventService.CreateEvent(request.Context(), eventCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (e EventControllerImpl) UpdateEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventUpdateRequest := web.EventUpdateRequest{}
	helper.ReadFromRequestBody(request, &eventUpdateRequest)

	eventId := params.ByName("eventId")
	id, err := strconv.Atoi(eventId)
	helper.PanicIfError(err)

	eventUpdateRequest.Event_id = id

	eventResponse := e.EventService.UpdateEvent(request.Context(), eventUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (e EventControllerImpl) DeleteEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventId := params.ByName("eventId")
	id, err := strconv.Atoi(eventId)
	helper.PanicIfError(err)

	e.EventService.DeleteEvent(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (e EventControllerImpl) FindEventById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventId := params.ByName("eventId")
	id, err := strconv.Atoi(eventId)
	helper.PanicIfError(err)

	eventResponse := e.EventService.FindEventById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (e EventControllerImpl) FindAllEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventResponse := e.EventService.FindAllEvent(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
