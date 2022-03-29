package usecase

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"teswir-go/internal/entity"
	"teswir-go/pkg/logger"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (u *useCase) socketRead(ctx context.Context, log logger.Interface, conn *websocket.Conn, quit chan interface{}) {
	/*	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			quit <- 1
			fmt.Println(err)
			return
		}

		var api entity.MongoApi
		err = json.Unmarshal(message, &api)
		if err != nil {
			fmt.Println(err)
			continue
		}

		user, err1 := u.ActionInfo(ctx, log, api.Token)
		if err1 != 0 {
			eMsg := "error in u.ActionInfo()"
			log.Error(eMsg, err1)
			continue
		}

		_, err = u.webAPI.ApiMongoUserGetByID(ctx, user.ID)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				err1 := u.webAPI.ApiMongoUserAdd(ctx, *user)
				if err1 != nil {
					eMsg := "error in u.webAPI.ApiMongoUserAdd"
					log.Error(eMsg, err1)
					continue
				}
			} else {
				eMsg := "error in u.webAPI.ApiMongoUserGetByID"
				log.Error(eMsg, err)
				continue
			}
		}

		if api.Api == "user/list" {
			users, err := u.webAPI.ApiMongoUserList(ctx)
			if err != nil {
				fmt.Println(err)
				continue
			}

			b, err1 := json.Marshal(users)
			if err1 != nil {
				fmt.Println(err1)
				continue
			}

			err = conn.WriteMessage(1, b)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if api.Api == "send-message" {

		}
	}*/
}

var m = make(map[uuid.UUID]*websocket.Conn)

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

		delete(m, actionInfo.ID)
	}()

	fmt.Println("NewConnection:", actionInfo.Username)

	user, err1 := u.webAPI.ApiMongoUserGetByID(ctx, actionInfo.ID)
	if err1 != nil && err1 != mongo.ErrNoDocuments {
		eMsg := "error in u.webAPI.ApiMongoUserGetByID"
		log.Error(eMsg, err1)
		return
	}

	if user != nil {
		err1 = u.webAPI.ApiMongoUserDeleteByID(ctx, actionInfo.ID)
		if err1 != nil {
			eMsg := "error in u.webAPI.ApiMongoUserDeleteByID"
			log.Error(eMsg)
			return
		}
	}

	err = u.webAPI.ApiMongoUserAdd(ctx, actionInfo)
	if err != nil {
		fmt.Println("error in apiMongoUserAdd")
		return
	}

	m[actionInfo.ID] = conn

	go u.webAPI.SocketRead(ctx, conn, m, quit)

	for {
		select {
		case <-quit:
			return
		}
	}
}
