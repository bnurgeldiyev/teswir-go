package entity

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
	"teswir-go/pkg/logger"
	"time"
)

type UserRole string

const (
	UserRoleAdmin  UserRole = "ADMIN"
	UserRoleClient UserRole = "CLIENT"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	UserRole  UserRole  `json:"user_role"`
	CreateTS  time.Time `json:"create_ts"`
	UpdateTS  time.Time `json:"update_ts"`
}

func NewUserEntity(log logger.Interface, id, username, firstname, lastname, userRole string) (item *User, err error) {

	m := make(map[string]string)
	m["id"] = id
	m["username"] = username
	m["firstname"] = firstname
	m["lastname"] = lastname
	m["userRole"] = userRole

	names, err := VerifyMinLen(m)
	if err != nil {
		eMsg := fmt.Sprintf("FormValue's <%s> is empty", names)
		log.Error(eMsg, err)
		return
	}

	realID, err := ConvertStringToUUID(id)
	if err != nil {
		eMsg := "error in convertStringToUUID(id)"
		log.Error(eMsg, err)
		return
	}

	role, err := ConvertStringToUserRole(userRole)
	if err != nil {
		eMsg := "error in convertStringToUserRole"
		log.Error(eMsg, err)
		return
	}

	item = &User{
		ID:        realID,
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		UserRole:  role,
	}

	return
}

func ConvertStringToUUID(str string) (id uuid.UUID, err error) {
	id, err = uuid.FromString(str)
	return
}

func ConvertStringToUserRole(str string) (role UserRole, err error) {

	roles := []UserRole{
		UserRoleAdmin,
		UserRoleClient,
	}

	str = strings.ToUpper(str)

	for _, r := range roles {
		if str == string(r) {
			role = r
			return
		}
	}

	err = errors.New("invalid user_role")

	return
}

func VerifyMinLen(m map[string]string) (names string, err error) {

	for k, v := range m {
		if len(v) == 0 {
			err = errors.New("len == 0")
			names += k
			names += " "
		}
	}

	return
}
