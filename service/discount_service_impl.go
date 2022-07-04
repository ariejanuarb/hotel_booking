package service

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator"
)

type DiscountServiceImpl struct {
	DiscountRepository repository.DiscountRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewDiscountService(discountRepository repository.DiscountRepository, db *sql.DB, validate *validator.Validate) DiscountService {
	return &DiscountServiceImpl{
		DiscountRepository: discountRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *DiscountServiceImpl) CreateDiscount(ctx context.Context, request web.DiscountCreateRequest) web.DiscountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount := domain.Discount{
		Discount_amount:  request.Discount_amount,
		Discount_request: request.Discount_request,
		Discount_status:  request.Discount_status,
		Invoice_id:       request.Invoice_id,
	}

	discount = service.DiscountRepository.SaveDiscount(ctx, tx, discount)
	return helper.ToDiscountResponse(discount)
}

func (service *DiscountServiceImpl) UpdateDiscount(ctx context.Context, request web.DiscountUpdateRequest) web.DiscountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount, err := service.DiscountRepository.FindDiscountById(ctx, tx, request.Discount_id)
	if err != nil {
		helper.PanicIfError(err)
	}
	discount.Discount_request = request.Discount_request
	discount.Discount_status = request.Discount_status
	discount.Discount_amount = request.Discount_amount

	discount = service.DiscountRepository.UpdateDiscount(ctx, tx, discount)
	return helper.ToDiscountResponse(discount)
}

func (service *DiscountServiceImpl) DeleteDiscount(ctx context.Context, discountId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount, err := service.DiscountRepository.FindDiscountById(ctx, tx, discountId)
	if err != nil {
		helper.PanicIfError(err)
	}

	service.DiscountRepository.DeleteDiscount(ctx, tx, discount)
}

func (service *DiscountServiceImpl) FindDiscountById(ctx context.Context, discountId int) web.DiscountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount, err := service.DiscountRepository.FindDiscountById(ctx, tx, discountId)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper.ToDiscountResponse(discount)
}

func (service DiscountServiceImpl) FindAllDiscount(ctx context.Context) []web.DiscountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discounts := service.DiscountRepository.FindAllDiscount(ctx, tx)

	return helper.ToDiscountResponses(discounts)
}
