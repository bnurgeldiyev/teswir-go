package usecase

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

func (u *useCase) UserAuth(ctx context.Context, log logger.Interface, username, password string) (item *entity.UserAuth, errCode int) {

	auth, err := u.webAPI.ApiAuth(ctx, username, password)
	if err != nil {
		eMsg := "error in u.webApi.Auth()"
		log.Error(eMsg, err)
		errCode = http.StatusUnauthorized
		return
	}

	item = &entity.UserAuth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}

	return
}

func (u *useCase) UserAdd(ctx context.Context, log logger.Interface, r *entity.User) (errCode int) {

	user, err1 := u.repo.UserGetByUsername(ctx, r.Username)
	if err1 != nil {
		eMsg := "error in u.repo.UserGetByUsername"
		log.Error(eMsg, err1)
		errCode = http.StatusInternalServerError
		return
	}

	if user != nil {
		eMsg := fmt.Sprintf("User with username=<%s> already exists", r.Username)
		log.Error(eMsg)
		errCode = http.StatusConflict
		return
	}

	err := u.repo.UserAdd(ctx, r)
	if err != nil {
		eMsg := "error in u.repo.UserAdd()"
		log.Error(eMsg, err)
		errCode = http.StatusInternalServerError
		return
	}

	return
}

func (u *useCase) UserGetByID(ctx context.Context, log logger.Interface, id uuid.UUID) (item *entity.User, errCode int) {

	user, err := u.repo.UserGetByID(ctx, id)
	if err != nil {
		eMsg := "error in u.repo.UserGetByID()"
		log.Error(eMsg, err)
		errCode = http.StatusInternalServerError
		return
	}

	if user == nil {
		eMsg := fmt.Sprintf("User with id=<%s> not found", id)
		log.Error(eMsg)
		errCode = http.StatusNotFound
		return
	}

	item = &entity.User{
		ID:        user.ID,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		UserRole:  user.UserRole,
		CreateTS:  user.CreateTS,
		UpdateTS:  user.UpdateTS,
	}

	return
}
