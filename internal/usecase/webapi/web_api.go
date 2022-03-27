package webapi

import (
	"go.mongodb.org/mongo-driver/mongo"
	"teswir-go/pkg/auth"
)

type WebAPI struct {
	auth       *auth.GrpcClient
	collection *mongo.Collection
}

func NewWebAPI(auth *auth.GrpcClient, collection *mongo.Collection) *WebAPI {
	return &WebAPI{
		auth:       auth,
		collection: collection,
	}
}
