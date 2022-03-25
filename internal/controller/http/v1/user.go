package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"teswir-go/internal/entity"
	"teswir-go/internal/usecase"
	"teswir-go/pkg/logger"
)

type userRoutes struct {
	uCase usecase.UseCase
	log   logger.Interface
}

func newUserRoutes(handler *mux.Router, uCase usecase.UseCase, l logger.Interface) {
	r := &userRoutes{uCase, l}

	handler.HandleFunc("/api/v1/user/{id}/get", r.userGetByID).Methods("GET")
	handler.HandleFunc("/api/v1/user/add", r.userAdd).Methods("POST")
	handler.HandleFunc("/api/v1/user/auth", r.userAuth).Methods("POST")
}

func (u *userRoutes) userAuth(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) == 0 {
		eMsg := "FormValue username is empty"
		u.log.Error(eMsg)
		errCode := http.StatusBadRequest
		SendResponse(w, nil, errCode)
		return
	}

	if len(password) == 0 {
		eMsg := "FormValue password is empty"
		u.log.Error(eMsg)
		errCode := http.StatusBadRequest
		SendResponse(w, nil, errCode)
		return
	}

	data, errCode := u.uCase.UserAuth(r.Context(), u.log, username, password)
	SendResponse(w, data, errCode)
}

func (u *userRoutes) userGetByID(w http.ResponseWriter, r *http.Request) {
	id, err1 := entity.ConvertStringToUUID(mux.Vars(r)["id"])
	if err1 != nil {
		eMsg := "error in entity.ConvertStringToUUID()"
		u.log.Error(eMsg, err1)
		errCode := http.StatusBadRequest
		SendResponse(w, nil, errCode)
		return
	}

	data, errCode := u.uCase.UserGetByID(r.Context(), u.log, id)
	SendResponse(w, data, errCode)
}

func (u *userRoutes) userAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := getTokenFromHeader(r)
	if err != nil {
		eMsg := "error in getTokenFromHeader()"
		u.log.Error(eMsg, err)
		errCode := http.StatusUnauthorized
		SendResponse(w, nil, errCode)
		return
	}

	_, eCode := u.uCase.ActionInfo(ctx, u.log, token)
	if eCode != 0 {
		eMsg := "error in p.uCase.ActionInfo()"
		u.log.Error(eMsg)
		errCode := http.StatusUnauthorized
		SendResponse(w, nil, errCode)
		return
	}

	user := new(entity.User)

	role, err := entity.ConvertStringToUserRole(r.FormValue("user_role"))
	if err != nil {
		eMsg := "error in entity.ConvertStringToUserRole(user_role)"
		u.log.Error(eMsg, err)
		errCode := http.StatusBadRequest
		SendResponse(w, nil, errCode)
		return
	}

	m := make(map[string]string)
	m["username"] = r.FormValue("username")
	m["firstname"] = r.FormValue("firstname")
	m["lastname"] = r.FormValue("lastname")

	names, err := entity.VerifyMinLen(m)
	if err != nil {
		eMsg := fmt.Sprintf("FormValue's <%s> is empty", names)
		u.log.Error(eMsg, err)
		errCode := http.StatusBadRequest
		SendResponse(w, nil, errCode)
		return
	}

	user.Username = m["username"]
	user.Firstname = m["firstname"]
	user.Lastname = m["lastname"]
	user.UserRole = role

	errCode := u.uCase.UserAdd(ctx, u.log, user)
	SendResponse(w, nil, errCode)
}
