package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"hotel_booking/exception"
	"hotel_booking/helper"
	"hotel_booking/model/domain"
	"hotel_booking/model/web"
	"hotel_booking/repository"
)

type MeetingRoomServiceImpl struct {
	MeetingRoomRepository repository.MeetingRoomRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewMeetingRoomService(meetingRoomRepository repository.MeetingRoomRepository, DB *sql.DB, validate *validator.Validate) MeetingRoomService {
	return &MeetingRoomServiceImpl{
		MeetingRoomRepository: meetingRoomRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *MeetingRoomServiceImpl) Create(ctx context.Context, request web.MeetingRoomCreateRequest) web.MeetingRoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom := domain.MeetingRoom{
		Name:         request.Name,
		Capacity:     request.Capacity,
		Facility:     request.Facility,
		PricePerHour: request.PricePerHour,
		PricePerDay:  request.PricePerDay,
		DailyRevenue: request.DailyRevenue,
	}

	meetingRoom = service.MeetingRoomRepository.Save(ctx, tx, meetingRoom)

	return helper.ToMeetingRoomResponse(meetingRoom)
}

func (service *MeetingRoomServiceImpl) Update(ctx context.Context, request web.MeetingRoomUpdateRequest) web.MeetingRoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom, err := service.MeetingRoomRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	meetingRoom.Name = request.Name
	meetingRoom.Capacity = request.Capacity
	meetingRoom.Facility = request.Facility
	meetingRoom.PricePerHour = request.PricePerHour
	meetingRoom.PricePerDay = request.PricePerDay
	meetingRoom.DailyRevenue = request.DailyRevenue

	meetingRoom = service.MeetingRoomRepository.Update(ctx, tx, meetingRoom)

	return helper.ToMeetingRoomResponse(meetingRoom)
}

func (service *MeetingRoomServiceImpl) Delete(ctx context.Context, meetingRoomId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom, err := service.MeetingRoomRepository.FindById(ctx, tx, meetingRoomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MeetingRoomRepository.Delete(ctx, tx, meetingRoom)
}

func (service *MeetingRoomServiceImpl) FindById(ctx context.Context, meetingRoomId int) web.MeetingRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRoom, err := service.MeetingRoomRepository.FindById(ctx, tx, meetingRoomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMeetingRoomResponse(meetingRoom)
}

func (service *MeetingRoomServiceImpl) FindAll(ctx context.Context) []web.MeetingRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meetingRooms := service.MeetingRoomRepository.FindAll(ctx, tx)

	return helper.ToMeetingRoomResponses(meetingRooms)
}
