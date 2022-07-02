package helper

import (
	"hotel_booking/model/domain"
	"hotel_booking/model/web"
)

func ToFloorResponses(floors []domain.Floor) []web.FloorResponse {
	var floorResponses []web.FloorResponse
	for _, floor := range floors {
		floorResponses = append(floorResponses, ToFloorResponse(floor))
	}
	return floorResponses
}

func ToFloorResponse(floor domain.Floor) web.FloorResponse {
	return web.FloorResponse{
		Id:       floor.Id,
		Number:   floor.Number,
		Capacity: floor.Capacity,
	}
}

func ToMeetingRoomResponses(meetingRooms []domain.MeetingRoom) []web.MeetingRoomResponse {
	var meetingRoomResponses []web.MeetingRoomResponse
	for _, meetingRoom := range meetingRooms {
		meetingRoomResponses = append(meetingRoomResponses, ToMeetingRoomResponse(meetingRoom))
	}
	return meetingRoomResponses
}

func ToMeetingRoomResponse(meetingRoom domain.MeetingRoom) web.MeetingRoomResponse {
	return web.MeetingRoomResponse{
		Id:           meetingRoom.Id,
		Name:         meetingRoom.Name,
		Capacity:     meetingRoom.Capacity,
		Facility:     meetingRoom.Facility,
		PricePerHour: meetingRoom.PricePerHour,
		PricePerDay:  meetingRoom.PricePerDay,
		DailyRevenue: meetingRoom.DailyRevenue,
	}
}

func ToRolesResponses(roless []domain.Roles) []web.RolesResponse {
	var rolesResponses []web.RolesResponse
	for _, roles := range roless {
		rolesResponses = append(rolesResponses, ToRolesResponse(roles))
	}
	return rolesResponses
}

func ToRolesResponse(roles domain.Roles) web.RolesResponse {
	return web.RolesResponse{
		Id:   roles.Id,
		Name: roles.Name,
	}
}
