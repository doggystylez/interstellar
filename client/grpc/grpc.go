package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(endpoint string, timeout int) Client {
	return Client{Endpoint: endpoint, Timeout: timeout, Ctx: context.Background()}
}

func (c *Client) Open() error {
	var err error
	c.Ctx, c.CtxCancel = context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	c.Conn, err = grpc.Dial(
		c.Endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return err
}

func (c *Client) Close() {
	c.CtxCancel()
	if err := c.Conn.Close(); err != nil {
		panic(err)
	}
}
