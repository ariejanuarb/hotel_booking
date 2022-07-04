package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type DiscountRepositoryImpl struct {
}

func NewDiscountRepository() DiscountRepository {
	return &DiscountRepositoryImpl{}
}

func (repository *DiscountRepositoryImpl) SaveDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount {
	SQL := "insert into discount(discount_request, discount_status, discount_amount) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, discount.Discount_request, discount.Discount_status, discount.Discount_amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	discount.Discount_id = int(id)
	return discount
}

func (repository *DiscountRepositoryImpl) UpdateDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount {
	SQL := "update discount set discount_status = ?, discount_amount = ? where discount_id = ?"
	_, err := tx.ExecContext(ctx, SQL, discount.Discount_status, discount.Discount_amount, discount.Discount_id)
	helper.PanicIfError(err)

	return discount
}

func (repository *DiscountRepositoryImpl) DeleteDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) {
	SQL := "delete from discount where discount_id = ?"
	_, err := tx.ExecContext(ctx, SQL, discount.Discount_id)
	helper.PanicIfError(err)
}

func (repository *DiscountRepositoryImpl) FindDiscountById(ctx context.Context, tx *sql.Tx, discountId int) (domain.Discount, error) {
	SQL := "select discount_id, discount_request, discount_status, discount_amount from discount where discount_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, discountId)
	helper.PanicIfError(err)
	defer rows.Close()

	discounts := domain.Discount{}

	if rows.Next() {
		err := rows.Scan(&discounts.Discount_id, &discounts.Discount_request, &discounts.Discount_status, &discounts.Discount_amount)
		helper.PanicIfError(err)
		defer rows.Close()
		return discounts, nil
	} else {
		return discounts, errors.New("discount is not found")
	}
}

func (repository *DiscountRepositoryImpl) FindAllDiscount(ctx context.Context, tx *sql.Tx) []domain.Discount {
	SQL := "select discount_id, discount_request, discount_status, discount_amount from discount"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var discount []domain.Discount
	for rows.Next() {
		discounts := domain.Discount{}
		err := rows.Scan(&discounts.Discount_id, &discounts.Discount_request, &discounts.Discount_status, &discounts.Discount_amount)
		helper.PanicIfError(err)
		discount = append(discount, discounts)
	}
	return discount
}
