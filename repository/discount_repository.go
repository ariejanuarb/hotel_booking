package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type DiscountRepository interface {
	SaveDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount
	UpdateDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount
	DeleteDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount)
	FindDiscountById(ctx context.Context, tx *sql.Tx, discountId int) (domain.Discount, error)
	FindAllDiscount(ctx context.Context, tx *sql.Tx) []domain.Discount
}
