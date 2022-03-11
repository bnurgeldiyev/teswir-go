package usecase

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

type UserUseCase struct {
	repo   UserRepo
	webAPI UserWebAPI
}

func NewUserUseCase(r UserRepo, w UserWebAPI) *UserUseCase {
	return &UserUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (u *UserUseCase) UserAdd(ctx context.Context, r *entity.User, log logger.Interface) (err error) {

	user, err1 := u.repo.UserGetByUsername(ctx, r.Username)
	if err1 != nil {
		
	}

	if user == nil {

	}

	err = u.repo.UserAdd(ctx, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (u *UserUseCase) UserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error) {

	item, err = u.repo.UserGetByID(ctx, id)
	if err != nil {
		//err = v1.ErrInternalServerError
		return
	}

	return
}
