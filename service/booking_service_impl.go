package service

import (
	"booking-hotel/exception"
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator"
)

type BookingServiceImpl struct {
	BookingRepository repository.BookingRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewBookingService(bookingRepository repository.BookingRepository, db *sql.DB, validate *validator.Validate) *BookingServiceImpl {
	return &BookingServiceImpl{
		BookingRepository: bookingRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *BookingServiceImpl) Create(ctx context.Context, request web.BookingCreateRequest) web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking := domain.Booking{
		Status:      request.Status,
		Pic_name:    request.Pic_name,
		Pic_Contact: request.Pic_Contact,
		Invoice:     request.Invoice,
		Event_start: request.Event_start,
		Event_end:   request.Event_end,
		Room_id:     request.Room_id,
	}
	booking = service.BookingRepository.Save(ctx, tx, booking)

	return helper.ToBookingResponse(booking)
}

func (service *BookingServiceImpl) Update(ctx context.Context, request web.BookingUpdateRequest) web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	bookings := helper.ToBooking(bookingResponse)
	bookings.Status = request.Status

	bookings = service.BookingRepository.Update(ctx, tx, bookings)

	return helper.ToBookingResponse(bookings)
}

func (service *BookingServiceImpl) FindById(ctx context.Context, bookingId int) web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking, err := service.BookingRepository.FindById(ctx, tx, bookingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return booking
}

func (service *BookingServiceImpl) FindAll(ctx context.Context) []web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking := service.BookingRepository.FindAll(ctx, tx)
	return booking
}
