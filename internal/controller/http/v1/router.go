package v1

import (
	"github.com/gorilla/mux"
	"teswir-go/internal/usecase"
	"teswir-go/pkg/logger"
)

func NewUserRouter(handler *mux.Router, l logger.Interface, user usecase.User) {
	newUserRoutes(handler, user, l)
}

func NewProductRouter(handler *mux.Router, l logger.Interface, product usecase.Product) {
	newProductRoutes(handler, product, l)
}
