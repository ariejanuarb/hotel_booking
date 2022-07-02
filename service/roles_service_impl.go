package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"hotel_booking/exception"
	"hotel_booking/helper"
	"hotel_booking/model/domain"
	"hotel_booking/model/web"
	"hotel_booking/repository"
)

type RolesServiceImpl struct {
	RolesRepository repository.RolesRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewRolesService(rolesRepository repository.RolesRepository, DB *sql.DB, validate *validator.Validate) RolesService {
	return &RolesServiceImpl{
		RolesRepository: rolesRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *RolesServiceImpl) Create(ctx context.Context, request web.RolesCreateRequest) web.RolesResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles := domain.Roles{
		Name: request.Name,
	}

	roles = service.RolesRepository.Save(ctx, tx, roles)

	return helper.ToRolesResponse(roles)
}

func (service *RolesServiceImpl) Update(ctx context.Context, request web.RolesUpdateRequest) web.RolesResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles, err := service.RolesRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	roles.Name = request.Name

	roles = service.RolesRepository.Update(ctx, tx, roles)

	return helper.ToRolesResponse(roles)
}

func (service *RolesServiceImpl) Delete(ctx context.Context, rolesId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles, err := service.RolesRepository.FindById(ctx, tx, rolesId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RolesRepository.Delete(ctx, tx, roles)
}

func (service *RolesServiceImpl) FindById(ctx context.Context, rolesId int) web.RolesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roles, err := service.RolesRepository.FindById(ctx, tx, rolesId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRolesResponse(roles)
}

func (service *RolesServiceImpl) FindAll(ctx context.Context) []web.RolesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roless := service.RolesRepository.FindAll(ctx, tx)

	return helper.ToRolesResponses(roless)
}
