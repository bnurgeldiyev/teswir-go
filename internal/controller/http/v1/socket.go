package v1

import (
	"github.com/gorilla/mux"
	"net/http"
	"teswir-go/internal/usecase"
	"teswir-go/pkg/logger"
)

type socketRoutes struct {
	uCase usecase.UseCase
	log   logger.Interface
}

func newSocketRoutes(handler *mux.Router, uCase usecase.UseCase, l logger.Interface) {
	r := &socketRoutes{uCase, l}

	handler.HandleFunc("/api/v1/socket", r.socket)
}

func (s *socketRoutes) socket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := getTokenFromHeader(r)
	if err != nil {
		eMsg := "error in getTokenFromHeader()"
		s.log.Error(eMsg, err)
		errCode := http.StatusUnauthorized
		SendResponse(w, nil, errCode)
		return
	}

	actionInfo, eCode := s.uCase.ActionInfo(ctx, s.log, token)
	if eCode != 0 {
		eMsg := "error in p.uCase.ActionInfo()"
		s.log.Error(eMsg)
		errCode := http.StatusUnauthorized
		SendResponse(w, nil, errCode)
		return
	}

	s.uCase.Socket(ctx, s.log, actionInfo, w, r)
}
