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
	"fmt"
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

func (service *BookingServiceImpl) Create(ctx context.Context, request *web.BookingCreateRequest) (*web.BookingResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	startTime := request.Event_start
	endTime := request.Event_end
	roomId := request.Room_id

	var price int
	var priceHour string
	var priceDay string
	var days int

	if startTime > time.Now().Format("2006-01-02 15:04:05") {
		checkSchadules := service.BookingRepository.FindAll(ctx, tx)
		for _, checkSchadule := range checkSchadules {
			if roomId == checkSchadule.Room_id {
				if checkSchadule.Status == "Booked" {
					if startTime >= checkSchadule.Event_start {
						if startTime <= checkSchadule.Event_end {
							return nil, errors.New("Room is Full")
						}
					} else {
						fmt.Println("Benar")
						if endTime >= checkSchadule.Event_start {
							return nil, errors.New("Room is Full")
						}
					}
				}
			}
			priceHour = checkSchadule.Price_per_hour
			priceDay = checkSchadule.Price_per_day
		}
		layout := "2006-01-02T15:04:05Z"

		timeStart, _ := time.Parse(layout, startTime)
		timeEnd, _ := time.Parse(layout, endTime)
		diff := timeEnd.Sub(timeStart).Hours()

		priceHours, _ := strconv.Atoi(priceHour)
		priceDays, _ := strconv.Atoi(priceDay)
		day := time.Hour.Hours() * 24

		if int(diff)%24 == 0 {
			days = int(diff / day)
		}

		if int(diff)%24 == 0 {
			price = priceDays * days
		} else if diff < day {
			price = priceHours * int(diff)
		}
	}

	booking := &domain.Booking{
		Status:             "Booked",
		Room_id:            request.Room_id,
		Pic_name:           request.Pic_name,
		Pic_Contact:        request.Pic_Contact,
		Event_start:        request.Event_start,
		Event_end:          request.Event_end,
		Invoice_number:     "INV-" + strconv.Itoa(request.Room_id) + "-" + time.Now().Format("20060402150405"),
		Invoice_subtotal:   priceHour + " or " + priceDay,
		Invoice_grandtotal: strconv.Itoa(price),
		Discount_request:   "Null",
		Discount_amount:    "0,00",
		Created_at:         time.Now(),
	}

	booking = service.BookingRepository.Save(ctx, tx, booking)

	return helper.ToBookingResponse(booking), nil
}

func (service *BookingServiceImpl) UpdateStatus(ctx context.Context, request *web.UpdateRequest) *web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	bookings := helper.ToBooking(*bookingResponse)
	bookings.Status = "Canceled"
	bookings.Updated_at = time.Now()

	bookings = service.BookingRepository.UpdateStatus(ctx, tx, bookings)

	return helper.ToBookingResponse(bookings)
}

func (service *BookingServiceImpl) UpdateDiscount(ctx context.Context, request *web.UpdateDiscount) *web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	bookings := helper.ToBooking(*bookingResponse)
	bookings.Discount_request = "Pending"
	bookings.Updated_at = time.Now()

	bookings = service.BookingRepository.UpdateDiscount(ctx, tx, bookings)

	return helper.ToBookingResponse(bookings)
}

func (service *BookingServiceImpl) ResponseDiscount(ctx context.Context, request *web.ResponseDiscount) *web.BookingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	bookings := helper.ToBooking(*bookingResponse)
	bookings.Discount_request = request.Discount_request
	bookings.Discount_amount = request.Discount_amount
	bookings.Updated_at = time.Now()

	bookings = service.BookingRepository.ResponseDiscount(ctx, tx, bookings)

	return helper.ToBookingResponse(bookings)
}

func (service *BookingServiceImpl) FindAllDiscount(ctx context.Context) []*web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount := service.BookingRepository.FindAllDiscount(ctx, tx)
	return discount
}

func (service *BookingServiceImpl) FindById(ctx context.Context, bookingId int) *web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking, err := service.BookingRepository.FindById(ctx, tx, bookingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return booking
}

func (service *BookingServiceImpl) FindAll(ctx context.Context) []*web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking := service.BookingRepository.FindAll(ctx, tx)
	return booking
}
