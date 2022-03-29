package usecase

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
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

		UserAdd(ctx context.Context, log logger.Interface, r *entity.User, password string) (errCode int)
		UserGetByID(ctx context.Context, log logger.Interface, id uuid.UUID) (item *entity.User, errCode int)
		UserAuth(ctx context.Context, log logger.Interface, username, password string) (item *entity.UserAuth, errCode int)
		UserList(ctx context.Context, log logger.Interface) (item []*entity.User, errCode int)

		Socket(ctx context.Context, log logger.Interface, actionInfo *entity.User, w http.ResponseWriter, r *http.Request)
		socketRead(ctx context.Context, log logger.Interface, conn *websocket.Conn, quit chan interface{})
	}

	Repo interface {
		RepoUserAdd(ctx context.Context, r *entity.User) (err error)
		RepoUserGetByUsername(ctx context.Context, username string) (item *entity.User, err error)
		RepoUserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error)
		RepoUserList(ctx context.Context) (item []*entity.User, err error)
	}

	WebAPI interface {
		ApiAuth(ctx context.Context, username, password string) (item *entity.UserAuth, err error)
		ApiVerifyToken(ctx context.Context, token string) (username string, err error)
		ApiCreate(ctx context.Context, username, password string) (err error)
		ApiDelete(ctx context.Context, username string) (err error)

		SocketRead(ctx context.Context, conn *websocket.Conn, m map[uuid.UUID]*websocket.Conn, quit chan interface{})

		ApiMongoUserAdd(ctx context.Context, user *entity.User) (err error)
		ApiMongoUserList(ctx context.Context) (item []*entity.User, err error)
		ApiMongoUserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error)
		ApiMongoUserDeleteByID(ctx context.Context, id uuid.UUID) (err error)
		ApiMongoUserDeleteAll(ctx context.Context) (err error)
	}
)
