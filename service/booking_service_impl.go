package service

import (
	"booking-hotel/exception"
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"strconv"
	"time"
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
		Status:             request.Status,
		Room_id:            request.Room_id,
		Pic_name:           request.Pic_name,
		Pic_Contact:        request.Pic_Contact,
		Event_start:        request.Event_start,
		Event_end:          request.Event_end,
		Invoice_number:     "INV-" + strconv.Itoa(request.Room_id) + "-" + time.Now().Format("20060402150405"),
		Invoice_subtotal:   "",
		Invoice_grandtotal: "",
		Discount_request:   "Pending",
		Discount_amount:    "0,00",
	}

	startTime := request.Event_start
	endTime := request.Event_end
	roomId := request.Room_id

	if startTime > time.Now().Format("2006-01-02 15:04:05") {
		checkSchadule, err := service.BookingRepository.GetEvent(ctx, tx, request.Room_id)
		if err == nil {
			if roomId == checkSchadule.Room_id {
				if startTime > checkSchadule.Event_start {
					if startTime < checkSchadule.Event_end {
						errors.New("Room is Full")
					}
				} else {
					if endTime > checkSchadule.Event_start {
						errors.New("Room is Full")
					}
				}
			}
		}
	} else {
		helper.PanicIfError(err)
	}

	booking = service.BookingRepository.Save(ctx, tx, booking)

	return helper.ToBookingResponse(booking)
}

func (service *BookingServiceImpl) UpdateStatus(ctx context.Context, request web.UpdateStatus) web.BookingResponse {
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

	bookings = service.BookingRepository.UpdateStatus(ctx, tx, bookings)

	return helper.ToBookingResponse(bookings)
}

func (service *BookingServiceImpl) UpdateDiscount(ctx context.Context, request web.UpdateDiscount) web.BookingResponse {
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
	bookings.Discount_request = request.Discount_request
	bookings.Discount_amount = request.Discount_amount

	bookings = service.BookingRepository.UpdateDiscount(ctx, tx, bookings)

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
