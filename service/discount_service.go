package service

import (
	"booking-hotel/model/web"
	"context"
)

type DiscountService interface {
	CreateDiscount(ctx context.Context, request web.DiscountCreateRequest) web.DiscountResponse
	UpdateDiscount(ctx context.Context, request web.DiscountUpdateRequest) web.DiscountResponse
	DeleteDiscount(ctx context.Context, discountId int)
	FindDiscountById(ctx context.Context, discount int) web.DiscountResponse
	FindAllDiscount(ctx context.Context) []web.DiscountResponse
}
