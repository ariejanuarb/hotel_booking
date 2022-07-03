package service

import (
	"booking-hotel/model/web"
	"context"
)

type EventService interface {
	CreateEvent(ctx context.Context, request web.EventCreateRequest) web.EventResponse
	UpdateEvent(ctx context.Context, request web.EventUpdateRequest) web.EventResponse
	DeleteEvent(ctx context.Context, eventId int)
	FindEventById(ctx context.Context, eventId int) web.EventResponse
	FindAllEvent(ctx context.Context) []web.EventResponse
}
