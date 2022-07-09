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
	router.GET("/api/booking/:bookingId", bookingController.FindById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
