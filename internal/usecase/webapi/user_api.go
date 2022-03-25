package webapi

import (
	"context"
	"teswir-go/internal/entity"
)

func (w *WebAPI) ApiAuth(ctx context.Context, username, password string) (item *entity.UserAuth, err error) {

	auth, err1 := w.auth.Auth(ctx, username, password)
	if err1 != nil {
		err = err1
	}

	if auth == nil {
		return
	}

	item = &entity.UserAuth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}

	return
}

func (w *WebAPI) ApiVerifyToken(ctx context.Context, token string) (username string, err error) {

	item, err1 := w.auth.VerifyToken(ctx, token)
	if err1 != nil {
		err = err1
	}

	if item == nil {
		return
	}

	username = item.Username

	return
}
