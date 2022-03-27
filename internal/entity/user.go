package entity

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"strings"
	"time"
	"unicode"
)

type UserRole string

const (
	UserRoleAdmin  UserRole = "ADMIN"
	UserRoleClient UserRole = "CLIENT"
)

type UserAuth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	UserRole  UserRole  `json:"user_role"`
	CreateTS  time.Time `json:"create_ts"`
	UpdateTS  time.Time `json:"update_ts"`
}

type MongoUser struct {
	ID        uuid.UUID       `json:"id"`
	Username  string          `json:"username"`
	Firstname string          `json:"firstname"`
	Lastname  string          `json:"lastname"`
	UserRole  UserRole        `json:"user_role"`
	Conn      *websocket.Conn `json:"conn"`
	CreateTS  time.Time       `json:"create_ts"`
	UpdateTS  time.Time       `json:"update_ts"`
}

type MongoApi struct {
	Api  string      `json:"api"`
	Data interface{} `json:"data"`
}

type MongoDataSendMessage struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}

func IsPasswordValid(pwd string) bool {

	var (
		hasMinLen = false
		hasNumber = false
	)

	if len(pwd) > 5 {
		hasMinLen = true
	}

	for _, char := range pwd {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	return hasMinLen && hasNumber
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
