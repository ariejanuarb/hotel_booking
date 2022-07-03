package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DiscountController interface {
	CreateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindDiscountById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
