package grpc

import (
	"context"
	"time"

	"github.com/doggystylez/interstellar/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Open(c *types.Client) (err error) {
	c.Ctx, c.CtxCancel = context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	c.Conn, err = grpc.Dial(
		c.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return
}

func Close(c *types.Client) {
	c.CtxCancel()
	c.Conn.Close()
}
