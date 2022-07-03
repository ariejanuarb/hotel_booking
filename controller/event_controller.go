package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type EventController interface {
	CreateEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindEventById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllEvent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
