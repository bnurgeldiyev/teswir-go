package usecase

import (
	"context"
	"github.com/gofrs/uuid"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

type useCase struct {
	repo   Repo
	webAPI WebAPI
}

func NewUseCase(r Repo, w WebAPI) *useCase {
	return &useCase{
		repo:   r,
		webAPI: w,
	}
}

type (
	UseCase interface {
		UserAdd(ctx context.Context, log logger.Interface, r *entity.User) (errCode int)
		UserGetByID(ctx context.Context, log logger.Interface, id uuid.UUID) (item *entity.User, errCode int)
		UserAuth(ctx context.Context, log logger.Interface, username, password string) (item *entity.UserAuth, errCode int)
	}

	Repo interface {
		UserAdd(ctx context.Context, r *entity.User) (err error)
		UserGetByUsername(ctx context.Context, username string) (item *entity.User, err error)
		UserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error)
	}

	WebAPI interface {
		ApiAuth(ctx context.Context, username, password string) (item *entity.UserAuth, err error)
	}
)
