package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type EventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

func (repository *EventRepositoryImpl) SaveEvent(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event {
	SQL := "insert into event(event_start, event_end) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, event.Event_start, event.Event_End)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	event.Event_id = int(id)
	return event
}

func (repository *EventRepositoryImpl) UpdateEvent(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event {
	SQL := "update event set event_start = ?, event_end = ? where event_id = ?"
	_, err := tx.ExecContext(ctx, SQL, event.Event_start, event.Event_End, event.Event_id)
	helper.PanicIfError(err)

	return event
}

func (repository *EventRepositoryImpl) DeleteEvent(ctx context.Context, tx *sql.Tx, event domain.Event) {
	SQL := "delete from event where event_id=?"
	_, err := tx.ExecContext(ctx, SQL, event.Event_id)
	helper.PanicIfError(err)
}

func (repository *EventRepositoryImpl) FindEventById(ctx context.Context, tx *sql.Tx, eventId int) (domain.Event, error) {
	SQL := "select event_id, event_start, event_end from event where event_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, eventId)
	helper.PanicIfError(err)
	defer rows.Close()

	events := domain.Event{}

	if rows.Next() {
		err := rows.Scan(&events.Event_id, &events.Event_start, &events.Event_End)
		helper.PanicIfError(err)
		defer rows.Close()
		return events, nil
	} else {
		return events, errors.New("event is not found")
	}
}

func (repository *EventRepositoryImpl) FindAllEvent(ctx context.Context, tx *sql.Tx) []domain.Event {
	SQL := "select event_id, event_start, event_end from event"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var event []domain.Event
	for rows.Next() {
		events := domain.Event{}
		err := rows.Scan(&events.Event_id, &events.Event_start, &events.Event_End)
		helper.PanicIfError(err)
		event = append(event, events)
	}
	return event
}
