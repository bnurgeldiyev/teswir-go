package v1

import (
	"github.com/gorilla/mux"
	"teswir-go/internal/usecase"
	"teswir-go/pkg/logger"
)

func NewRouter(handler *mux.Router, l logger.Interface, uCase usecase.UseCase) {
	newUserRoutes(handler, uCase, l)
}
