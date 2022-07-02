package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel_booking/helper"
	"hotel_booking/model/domain"
)

type RolesRepositoryImpl struct {
}

func NewRolesRepository() RolesRepository {
	return &RolesRepositoryImpl{}
}

func (repository *RolesRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, roles domain.Roles) domain.Roles {
	SQL := "insert into roles(role_name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, roles.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	roles.Id = int(id)
	return roles
}

func (repository *RolesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, roles domain.Roles) domain.Roles {
	SQL := "update roles set role_name = ? where roles_id = ?"
	_, err := tx.ExecContext(ctx, SQL, roles.Name, roles.Id)
	helper.PanicIfError(err)

	return roles
}

func (repository *RolesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, roles domain.Roles) {
	SQL := "delete from roles where roles_id = ?"
	_, err := tx.ExecContext(ctx, SQL, roles.Id)
	helper.PanicIfError(err)
}

func (repository *RolesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, rolesId int) (domain.Roles, error) {
	SQL := "select roles_id, roles_name from roles where roles_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, rolesId)
	helper.PanicIfError(err)
	defer rows.Close()

	roles := domain.Roles{}
	if rows.Next() {
		err := rows.Scan(&roles.Id, &roles.Name)
		helper.PanicIfError(err)
		return roles, nil
	} else {
		return roles, errors.New("roles is not found")
	}
}

func (repository *RolesRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Roles {
	SQL := "select roles_id, roles_name from roles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roless []domain.Roles
	for rows.Next() {
		roles := domain.Roles{}
		err := rows.Scan(&roles.Id, &roles.Name)
		helper.PanicIfError(err)
		roless = append(roless, roles)
	}
	return roless
}
