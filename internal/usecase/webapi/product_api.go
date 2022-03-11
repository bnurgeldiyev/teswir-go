package webapi

import (
	"context"
	"teswir-go/internal/entity"
)

type ProductWebAPI struct{}

func NewProductWebAPI() *ProductWebAPI {
	return &ProductWebAPI{}
}

func (p ProductWebAPI) ProductAdd(ctx context.Context, r *entity.Product) (err error) {
	return
}
