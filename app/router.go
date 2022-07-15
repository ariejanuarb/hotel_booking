package app

import (
	"booking-hotel/controller"
	"booking-hotel/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookingController controller.BookingController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/booking", bookingController.Create)
	router.GET("/api/booking", bookingController.FindAll)
	router.PUT("/api/bookingStatus/:bookingId", bookingController.UpdateStatus)
	router.PUT("/api/bookingDiscount/:bookingId", bookingController.UpdateDiscount)
	router.PUT("/api/responseDiscount/:bookingId", bookingController.ResponseDiscount)
	router.GET("/api/booking/:bookingId", bookingController.FindById)
	router.GET("/api/bookingDiscount", bookingController.FindAllDiscount)

	router.PanicHandler = exception.ErrorHandler

	return router
}
