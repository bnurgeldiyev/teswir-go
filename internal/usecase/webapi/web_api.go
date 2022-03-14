package webapi

import "teswir-go/pkg/auth"

type WebAPI struct {
	auth *auth.GrpcClient
}

func NewWebAPI(auth *auth.GrpcClient) *WebAPI {
	return &WebAPI{
		auth: auth,
	}
}
