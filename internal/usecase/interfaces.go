package usecase

import (
	"context"
	"github.com/gofrs/uuid"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

type (
	User interface {
		UserAdd(ctx context.Context, r *entity.User, log logger.Interface) (err error)
		UserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error)
	}

	UserRepo interface {
		UserAdd(ctx context.Context, r *entity.User) (err error)
		UserGetByUsername(ctx context.Context, username string) (item *entity.User, err error)
		UserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error)
	}

	UserWebAPI interface {
		UserAdd(ctx context.Context, r *entity.User) (err error)
	}
)

type (
	Product interface {
		ProductAdd(ctx context.Context, r *entity.Product) (err error)
	}

	ProductRepo interface {
		ProductAdd(ctx context.Context, r *entity.Product) (err error)
	}

	ProductWebAPI interface {
		ProductAdd(ctx context.Context, r *entity.Product) (err error)
	}
)
