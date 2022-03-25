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

func read(conn *websocket.Conn, quit chan interface{}) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			quit <- 1
			fmt.Println(err)
			return
		}

		fmt.Println(string(message))
	}
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
	}()

	fmt.Println("NewConnection:", r.RemoteAddr)

	go read(conn, quit)

	for {
		select {
		case <-quit:
			return
		}
	}
}
