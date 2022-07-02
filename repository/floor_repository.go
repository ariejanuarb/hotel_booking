package repository

import (
	"context"
	"database/sql"
	"hotel_booking/model/domain"
)

type FloorRepository interface {
	Save(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor
	Update(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor
	Delete(ctx context.Context, tx *sql.Tx, floor domain.Floor)
	FindById(ctx context.Context, tx *sql.Tx, floor int) (domain.Floor, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Floor
}
