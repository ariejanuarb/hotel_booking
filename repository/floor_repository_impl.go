package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel_booking/helper"
	"hotel_booking/model/domain"
)

type FloorRepositoryImpl struct {
}

func NewFloorRepository() FloorRepository {
	return &FloorRepositoryImpl{}
}

func (repository *FloorRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	SQL := "insert into floor(floor_number, floor_capacity) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, floor.Number, floor.Capacity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	floor.Id = int(id)
	return floor
}

func (repository *FloorRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	SQL := "update floor set floor_number = ?, floor_capacity = ? where floor_id = ?"
	_, err := tx.ExecContext(ctx, SQL, floor.Number, floor.Capacity, floor.Id)
	helper.PanicIfError(err)

	return floor
}

func (repository *FloorRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, floor domain.Floor) {
	SQL := "delete from floor where floor_id = ?"
	_, err := tx.ExecContext(ctx, SQL, floor.Id)
	helper.PanicIfError(err)
}

func (repository *FloorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, floorId int) (domain.Floor, error) {
	SQL := "select floor_id, floor_number, floor_capacity from floor where floor_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, floorId)
	helper.PanicIfError(err)
	defer rows.Close()

	floor := domain.Floor{}
	if rows.Next() {
		err := rows.Scan(&floor.Id, &floor.Number, &floor.Capacity)
		helper.PanicIfError(err)
		return floor, nil
	} else {
		return floor, errors.New("floor is not found")
	}
}

func (repository *FloorRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Floor {
	SQL := "select floor_id, floor_number, floor_capacity from floor"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var floors []domain.Floor
	for rows.Next() {
		floor := domain.Floor{}
		err := rows.Scan(&floor.Id, &floor.Number, &floor.Capacity)
		helper.PanicIfError(err)
		floors = append(floors, floor)
	}
	return floors
}
