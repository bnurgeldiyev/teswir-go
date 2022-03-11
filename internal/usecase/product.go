package usecase

import (
	"context"
	"teswir-go/internal/entity"
)

type ProductUseCase struct {
	repo   ProductRepo
	webAPI ProductWebAPI
}

func NewProductUseCase(r ProductRepo, w ProductWebAPI) *ProductUseCase {
	return &ProductUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (p *ProductUseCase) ProductAdd(ctx context.Context, r *entity.Product) (err error) {
	return
}
