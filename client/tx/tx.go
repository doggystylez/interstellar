package tx

import (
	"errors"
	"strconv"
	"time"

	"github.com/doggystylez/interstellar-proto/tx"
	"github.com/doggystylez/interstellar/client/grpc"
)

func AssembleAndBroadcast(msgs []MsgInfo, txInfo TxInfo, rpc grpc.Client) (TxResponse, error) {
	txBytes, err := SignFromPrivkey(makeMsgs(msgs), txInfo)
	if err != nil {
		return TxResponse{}, err
	}
	return broadcastTx(txBytes, rpc)
}

func broadcastTx(txBytes []byte, g grpc.Client) (TxResponse, error) {
	var grpcRes *tx.BroadcastTxResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	txClient := tx.NewServiceClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		grpcRes, err = txClient.BroadcastTx(
			g.Ctx,
			&tx.BroadcastTxRequest{
				Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
				TxBytes: txBytes,
			},
		)
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			return TxResponse{
				Code: &grpcRes.TxResponse.Code,
				Hash: &grpcRes.TxResponse.TxHash,
				Log:  &grpcRes.TxResponse.RawLog,
			}, nil
		}
	}
	return TxResponse{}, errors.New("failed sending tx after " + strconv.Itoa(g.Retries+1) + " attempts. last err: " + err.Error())
}
