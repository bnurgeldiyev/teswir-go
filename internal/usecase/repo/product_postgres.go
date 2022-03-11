package repo

import (
	"context"
	"teswir-go/internal/entity"
	"teswir-go/pkg/postgres"
)

type ProductRepo struct {
	*postgres.Postgres
}

func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

func (u ProductRepo) ProductAdd(ctx context.Context, r *entity.Product) (err error) {
	return
}
