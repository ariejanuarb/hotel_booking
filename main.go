package main

import (
	"booking-hotel/app"
	"booking-hotel/controller"
	"booking-hotel/helper"
	"booking-hotel/middleware"
	"booking-hotel/repository"
	"booking-hotel/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, db, validate)
	bookingController := controller.NewBookingController(bookingService)

	router := app.NewRouter(bookingController)

	server := http.Server{
		Addr:    "localhost:3080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
