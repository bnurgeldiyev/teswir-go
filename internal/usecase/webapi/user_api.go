package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"teswir-go/internal/entity"
)

func (w *WebAPI) ApiAuth(ctx context.Context, username, password string) (item *entity.UserAuth, err error) {

	auth, err1 := w.auth.Auth(ctx, username, password)
	if err1 != nil {
		err = err1
	}

	if auth == nil {
		return
	}

	item = &entity.UserAuth{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}

	return
}

func (w *WebAPI) ApiVerifyToken(ctx context.Context, token string) (username string, err error) {

	item, err1 := w.auth.VerifyToken(ctx, token)
	if err1 != nil {
		err = err1
	}

	if item == nil {
		return
	}

	username = item.Username

	return
}

func (w *WebAPI) ApiCreate(ctx context.Context, username, password string) (err error) {

	_, err1 := w.auth.Create(ctx, username, password)
	if err1 != nil {
		err = err1
	}

	return
}

func (w *WebAPI) ApiDelete(ctx context.Context, username string) (err error) {

	_, err1 := w.auth.Delete(ctx, username)
	if err1 != nil {
		err = err1
	}

	return
}

func (w *WebAPI) ApiMongoUserAdd(ctx context.Context, user *entity.User) (err error) {

	_, err = w.collection.InsertOne(ctx, user)

	return
}

func (w *WebAPI) ApiMongoUserList(ctx context.Context) (item []*entity.User, err error) {

	rows, err1 := w.collection.Find(ctx, bson.D{})
	if err1 != nil {
		err = err1
		return
	}

	defer func() {
		err := rows.Close(ctx)
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next(ctx) {
		var user *entity.User
		err := rows.Decode(&user)
		if err != nil {
			panic(err)
		}

		item = append(item, user)
	}

	return
}

func (w *WebAPI) ApiMongoUserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error) {

	filter := bson.D{{"id", id}}

	err = w.collection.FindOne(ctx, filter).Decode(&item)

	return
}

func (w *WebAPI) ApiMongoUserDeleteByID(ctx context.Context, id uuid.UUID) (err error) {

	filter := bson.D{{"id", id}}
	_, err = w.collection.DeleteOne(ctx, filter)

	return
}

func (w *WebAPI) ApiMongoUserDeleteAll(ctx context.Context) (err error) {

	_, err = w.collection.DeleteMany(ctx, bson.D{})

	return
}

func (w *WebAPI) SocketRead(ctx context.Context, conn *websocket.Conn, m map[uuid.UUID]*websocket.Conn, quit chan interface{}) {
	for {
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

		if api.Api == "user/list" {
			users, err := w.ApiMongoUserList(ctx)
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
			b, err := json.Marshal(api.Data)
			if err != nil {
				fmt.Println(err)
				continue
			}

			var sendMessage entity.MongoDataSendMessage
			err = json.Unmarshal(b, &sendMessage)
			if err != nil {
				fmt.Println(err)
				continue
			}

			c, ok := m[sendMessage.UserID]
			if !ok {
				fmt.Println("User not found")
				continue
			}

			err = c.WriteMessage(1, []byte(sendMessage.Message))
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
