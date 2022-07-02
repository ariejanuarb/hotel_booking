package repository

import (
	"context"
	"database/sql"
	"hotel_booking/model/domain"
)

type MeetingRoomRepository interface {
	Save(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom
	Update(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom) domain.MeetingRoom
	Delete(ctx context.Context, tx *sql.Tx, meetingRoom domain.MeetingRoom)
	FindById(ctx context.Context, tx *sql.Tx, meetingRoom int) (domain.MeetingRoom, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoom
}
