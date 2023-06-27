package types

import (
	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/doggystylez/interstellar/client/tx"
)

type InterstellarConfig struct {
	Path   string
	TxInfo tx.TxInfo
	Rpc    grpc.Client
}
