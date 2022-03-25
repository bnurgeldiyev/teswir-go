package auth

import (
	"context"
	"google.golang.org/grpc"
	"teswir-go/pkg/logger"
	"time"
)

type GrpcClient struct {
	AuthClient
}

func NewGrpcClient(url string, log logger.Interface) (client *GrpcClient, err error) {

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return
	}

	c := NewAuthClient(conn)

	go func() {
		ok := false
		for {
			request := &PingRequest{
				Ping: "ping",
			}

			_, err1 := c.Ping(context.Background(), request)
			if err1 != nil {
				eMsg := "error in c.Ping()"
				log.Error(eMsg, err1)
				ok = false
			} else {
				if !ok {
					log.Info("Connected to grpc server!")
					ok = true
				}
			}

			time.Sleep(3 * time.Second)
		}
	}()

	return &GrpcClient{
		c,
	}, nil
}

func (client *GrpcClient) Auth(ctx context.Context, username, password string) (item *UserAuthResponse, err error) {

	request := &UserAuthRequest{
		Username: username,
		Password: password,
	}

	item, err = client.UserAuth(ctx, request)

	return
}

func (client *GrpcClient) VerifyToken(ctx context.Context, token string) (item *UserAccessResponse, err error) {

	request := &UserAccessRequest{
		AccessToken: token,
	}

	item, err = client.UserAccess(ctx, request)

	return
}
