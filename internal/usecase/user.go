package usecase

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

func (u *useCase) ActionInfo(ctx context.Context, log logger.Interface, token string) (item *entity.User, errCode int) {

	username, err := u.webAPI.ApiVerifyToken(ctx, token)
	if err != nil {
		eMsg := "error in u.webAPI.ApiVerifyToken"
		log.Error(eMsg, err)
		errCode = http.StatusUnauthorized
		return
	}

	user, err := u.repo.RepoUserGetByUsername(ctx, username)
	if err != nil {
		eMsg := "error in u.repo.RepoUserGetByUsername"
		log.Error(eMsg, err)
		errCode = http.StatusInternalServerError
		return
	}

	if user == nil {
		eMsg := fmt.Sprintf("User with username=<%s> not found", username)
		log.Error(eMsg)
		errCode = http.StatusUnauthorized
		return
	}

	item = &entity.User{
		ID:        user.ID,
		Username:  username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		UserRole:  user.UserRole,
		CreateTS:  user.CreateTS,
		UpdateTS:  user.UpdateTS,
	}

	return
}

func (u *useCase) UserList(ctx context.Context, log logger.Interface) (item []*entity.User, errCode int) {

	users, err := u.repo.RepoUserList(ctx)
	if err != nil {
		eMsg := "error in u.repo.RepoUserList"
		log.Error(eMsg, err)
		errCode = http.StatusInternalServerError
		return
	}

	item = users

	return
}

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

func (u *useCase) UserAdd(ctx context.Context, log logger.Interface, r *entity.User, password string) (errCode int) {

	user, err1 := u.repo.RepoUserGetByUsername(ctx, r.Username)
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

	err := u.webAPI.ApiCreate(ctx, r.Username, password)
	if err != nil {
		eMsg := "error in u.webAPI.ApiCreate"
		log.Error(eMsg, err)
		errCode = http.StatusInternalServerError
		return
	}

	defer func() {
		if errCode != 0 {
			err := u.webAPI.ApiDelete(ctx, r.Username)
			if err != nil {
				log.Error("Error in u.webAPI.ApiDelete")
			}
		}
	}()

	err = u.repo.RepoUserAdd(ctx, r)
	if err != nil {
		eMsg := "error in u.repo.UserAdd()"
		log.Error(eMsg, err)
		errCode = http.StatusInternalServerError
		return
	}

	return
}

func (u *useCase) UserGetByID(ctx context.Context, log logger.Interface, id uuid.UUID) (item *entity.User, errCode int) {

	user, err := u.repo.RepoUserGetByID(ctx, id)
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
