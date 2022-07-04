package service

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator"
)

type EventServiceImpl struct {
	EventRepository repository.EventRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewEventService(eventRepository repository.EventRepository, db *sql.DB, validate *validator.Validate) EventService {
	return &EventServiceImpl{
		EventRepository: eventRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *EventServiceImpl) CreateEvent(ctx context.Context, request web.EventCreateRequest) web.EventResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	event := domain.Event{
		Event_start: request.Event_start,
		Event_End:   request.Event_end,
	}

	event = service.EventRepository.SaveEvent(ctx, tx, event)
	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) UpdateEvent(ctx context.Context, request web.EventUpdateRequest) web.EventResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	event, err := service.EventRepository.FindEventById(ctx, tx, request.Event_id)
	if err != nil {
		helper.PanicIfError(err)
	}
	event.Event_start = request.Event_start
	event.Event_End = request.Event_end

	event = service.EventRepository.UpdateEvent(ctx, tx, event)
	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) DeleteEvent(ctx context.Context, eventId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	event, err := service.EventRepository.FindEventById(ctx, tx, eventId)
	if err != nil {
		helper.PanicIfError(err)
	}

	service.EventRepository.DeleteEvent(ctx, tx, event)
}

func (service *EventServiceImpl) FindEventById(ctx context.Context, eventId int) web.EventResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	event, err := service.EventRepository.FindEventById(ctx, tx, eventId)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) FindAllEvent(ctx context.Context) []web.EventResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	events := service.EventRepository.FindAllEvent(ctx, tx)

	return helper.ToEventResponses(events)
}
