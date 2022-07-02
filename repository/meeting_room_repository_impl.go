package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel_booking/helper"
	"hotel_booking/model/domain"
)

type MeetingRoomRepositoryImpl struct {
}

func NewMeetingRoomRepository() MeetingRoomRepository {
	return &MeetingRoomRepositoryImpl{}
}

func (repository *MeetingRoomRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom {
	SQL := "insert into meeting_room(name_room, meeting_room_capacity, facility, price_per_hour, price_per_day, meeting_room_daily_revenue) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, meetingRoom.Name, meetingRoom.Capacity, meetingRoom.Facility, meetingRoom.PricePerHour, meetingRoom.PricePerDay, meetingRoom.DailyRevenue)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	meetingRoom.Id = int(id)
	return meetingRoom
}

func (repository *MeetingRoomRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom {
	SQL := "update meeting_room set name_room = ?, meeting_room_capacity = ?, facility = ?, price_per_hour = ?, price_per_day = ?, meeting_room_daily_revenue = ?  where floor_id = ?"
	_, err := tx.ExecContext(ctx, SQL, meetingRoom.Name, meetingRoom.Capacity, meetingRoom.Facility, meetingRoom.PricePerHour, meetingRoom.PricePerDay, meetingRoom.DailyRevenue, meetingRoom.Id)
	helper.PanicIfError(err)

	return meetingRoom
}

func (repository *MeetingRoomRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) {
	SQL := "delete from meeting_room where meeting_room_id = ?"
	_, err := tx.ExecContext(ctx, SQL, meetingRoom.Id)
	helper.PanicIfError(err)
}

func (repository *MeetingRoomRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, meetingRoomId int) (domain.MeetingRoom, error) {
	SQL := "select name_room, meeting_room_capacity, facility, price_per_hour, price_per_day, meeting_room_daily_revenue from floor where meeting_room_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, meetingRoomId)
	helper.PanicIfError(err)
	defer rows.Close()

	meetingRoom := domain.MeetingRoom{}
	if rows.Next() {
		err := rows.Scan(&meetingRoom.Name, &meetingRoom.Capacity, &meetingRoom.Facility, &meetingRoom.PricePerHour, &meetingRoom.PricePerDay, &meetingRoom.DailyRevenue, &meetingRoom.Id)
		helper.PanicIfError(err)
		return meetingRoom, nil
	} else {
		return meetingRoom, errors.New("meetingRoom is not found")
	}
}

func (repository *MeetingRoomRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoom {
	SQL := "select meeting_room_id, name_room, meeting_room_capacity, facility, price_per_hour, price_per_day, meeting_room_daily_revenue from meeting_room"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var meetingRooms []domain.MeetingRoom
	for rows.Next() {
		meetingRoom := domain.MeetingRoom{}
		err := rows.Scan(&meetingRoom.Id, &meetingRoom.Name, &meetingRoom.Capacity, &meetingRoom.Facility, &meetingRoom.PricePerHour, &meetingRoom.PricePerDay, &meetingRoom.DailyRevenue)
		helper.PanicIfError(err)
		meetingRooms = append(meetingRooms, meetingRoom)
	}
	return meetingRooms
}
