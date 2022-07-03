package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type EventRepository interface {
	SaveEvent(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	UpdateEvent(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	DeleteEvent(ctx context.Context, tx *sql.Tx, event domain.Event)
	FindEventById(ctx context.Context, tx *sql.Tx, eventId int) (domain.Event, error)
	FindAllEvent(ctx context.Context, tx *sql.Tx) []domain.Event
}
