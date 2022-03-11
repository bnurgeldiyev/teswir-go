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
	uCase usecase.User
	log   logger.Interface
}

func newUserRoutes(handler *mux.Router, u usecase.User, l logger.Interface) {
	r := &userRoutes{u, l}

	handler.HandleFunc("/api/v1/user/{id}/get", r.userGetByID).Methods("GET")
	handler.HandleFunc("/api/v1/user/add", r.userAdd).Methods("POST")
}

func (u *userRoutes) userGetByID(w http.ResponseWriter, r *http.Request) {
	id, err1 := entity.ConvertStringToUUID(mux.Vars(r)["id"])
	if err1 != nil {
		eMsg := "error in entity.ConvertStringToUUID()"
		u.log.Error(eMsg, err1)
		err := ErrBadRequest
		SendResponseByErrCode(w, err)
		return
	}

	data, err := u.uCase.UserGetByID(r.Context(), id)
	if err != nil {
		eMsg := "error in u.uCase.UserGetByID()"
		u.log.Error(eMsg, err)
		SendResponseByErrCode(w, err)
		return
	}

	if data == nil {
		eMsg := fmt.Sprintf("User with id=<%s> not found", id)
		u.log.Error(eMsg)
		SendResponseByErrCode(w, ErrNotFound)
		return
	}

	SendResponseOKWithData(w, data)
}

func (u *userRoutes) userAdd(w http.ResponseWriter, r *http.Request) {
	user := new(entity.User)

	role, err := entity.ConvertStringToUserRole(r.FormValue("user_role"))
	if err != nil {
		eMsg := "error in entity.ConvertStringToUserRole(user_role)"
		u.log.Error(eMsg, err)
		err = ErrBadRequest
		SendResponseByErrCode(w, err)
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
		err = ErrBadRequest
		SendResponseByErrCode(w, err)
		return
	}

	user.Username = m["username"]
	user.Firstname = m["firstname"]
	user.Lastname = m["lastname"]
	user.UserRole = role

	err = u.uCase.UserAdd(r.Context(), user, u.log)
	if err != nil {
		eMsg := "error in u.uCase.UserAdd()"
		u.log.Error(eMsg, err)
		SendResponseByErrCode(w, err)
		return
	}

	SendResponseOKWithData(w, nil)
}
