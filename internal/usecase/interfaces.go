package usecase

import (
	"context"
	"github.com/gofrs/uuid"
	"net/http"
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
		ActionInfo(ctx context.Context, log logger.Interface, token string) (item *entity.User, errCode int)

		UserAdd(ctx context.Context, log logger.Interface, r *entity.User) (errCode int)
		UserGetByID(ctx context.Context, log logger.Interface, id uuid.UUID) (item *entity.User, errCode int)
		UserAuth(ctx context.Context, log logger.Interface, username, password string) (item *entity.UserAuth, errCode int)

		Socket(ctx context.Context, log logger.Interface, actionInfo *entity.User, w http.ResponseWriter, r *http.Request)
	}

	Repo interface {
		RepoUserAdd(ctx context.Context, r *entity.User) (err error)
		RepoUserGetByUsername(ctx context.Context, username string) (item *entity.User, err error)
		RepoUserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error)
	}

	WebAPI interface {
		ApiAuth(ctx context.Context, username, password string) (item *entity.UserAuth, err error)
		ApiVerifyToken(ctx context.Context, token string) (username string, err error)
	}
)
