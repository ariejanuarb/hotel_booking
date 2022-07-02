package repository

import (
	"context"
	"database/sql"
	"hotel_booking/model/domain"
)

type RolesRepository interface {
	Save(ctx context.Context, tx *sql.Tx, roles domain.Roles) domain.Roles
	Update(ctx context.Context, tx *sql.Tx, roles domain.Roles) domain.Roles
	Delete(ctx context.Context, tx *sql.Tx, roles domain.Roles)
	FindById(ctx context.Context, tx *sql.Tx, roles int) (domain.Roles, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Roles
}
