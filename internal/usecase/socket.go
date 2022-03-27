package usecase

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (u *useCase) Socket(ctx context.Context, log logger.Interface, actionInfo *entity.User, w http.ResponseWriter, r *http.Request) {

	quit := make(chan interface{})

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Error during connection upgradation:", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("error in conn.Close()")
		}

		err = u.webAPI.ApiMongoUserDeleteByID(ctx, actionInfo.ID)
		if err != nil {
			log.Error("error in u.webAPI.ApiMongoUserDeleteByID")
		}
	}()

	fmt.Println("NewConnection:", actionInfo.Username)

	var mongoUser entity.User
	mongoUser.ID = actionInfo.ID
	mongoUser.Username = actionInfo.Username
	mongoUser.Firstname = actionInfo.Firstname
	mongoUser.Lastname = actionInfo.Lastname
	mongoUser.UserRole = actionInfo.UserRole
	mongoUser.CreateTS = actionInfo.CreateTS
	mongoUser.UpdateTS = actionInfo.UpdateTS

	err = u.webAPI.ApiMongoUserAdd(ctx, mongoUser)
	if err != nil {
		fmt.Println("error in apiMongoUserAdd")
		return
	}

	go u.webAPI.SocketRead(ctx, actionInfo.ID, conn, quit)

	for {
		select {
		case <-quit:
			return
		}
	}
}
