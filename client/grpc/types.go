package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type (
	Client struct {
		Endpoint  string
		Timeout   int
		Ctx       context.Context
		CtxCancel func()
		Conn      *grpc.ClientConn
	}
)
