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

	item = &entity.UserAuth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}

	return
}
