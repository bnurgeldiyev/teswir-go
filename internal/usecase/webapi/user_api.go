package webapi

import (
	"context"
	"teswir-go/internal/entity"
)

type UserWebAPI struct{}

func NewUserWebAPI() *UserWebAPI {
	return &UserWebAPI{}
}

func (u UserWebAPI) UserAdd(ctx context.Context, r *entity.User) (err error) {
	//TODO implement me
	panic("implement me")
}
