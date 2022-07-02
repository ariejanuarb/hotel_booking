package service

import (
	"context"
	"hotel_booking/model/web"
)

type MeetingRoomService interface {
	Create(ctx context.Context, request web.MeetingRoomCreateRequest) web.MeetingRoomResponse
	Update(ctx context.Context, request web.MeetingRoomUpdateRequest) web.MeetingRoomResponse
	Delete(ctx context.Context, meetingRoomId int)
	FindById(ctx context.Context, meetingRoomId int) web.MeetingRoomResponse
	FindAll(ctx context.Context) []web.MeetingRoomResponse
}
