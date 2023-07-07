package query

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/doggystylez/interstellar-proto/tx"
	"github.com/doggystylez/interstellar/client/grpc"
)

func Tx(hash *string, g grpc.Client) (*tx.GetTxResponse, error) {
	*hash = strings.ToUpper(*hash)
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := tx.NewServiceClient(g.Conn)
	var res *tx.GetTxResponse
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.GetTx(g.Ctx, &tx.GetTxRequest{Hash: *hash})
		if err != nil {
			if strings.Contains(err.Error(), *hash) && strings.Contains(err.Error(), "not found") {
				return &tx.GetTxResponse{}, TxNotFoundErr{message: "tx " + *hash + " not found"}
			}
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			return res, nil
		}
	}
	return &tx.GetTxResponse{}, RetryErr{retries: g.Retries, err: err}
}

func TxCode(hash string, g grpc.Client) (int, error) {
	txRes, err := Tx(&hash, g)
	if err != nil {
		return -1, err
	}
	return int(txRes.TxResponse.Code), nil
}

func AwaitTx(hash string, waitTime int, g grpc.Client) (*tx.GetTxResponse, error) {
	var (
		err   error
		txRes *tx.GetTxResponse
	)
	blockTime, timeNow := 6, time.Now()

	for tries := -1; tries < g.Retries; tries++ {
		txRes, err = Tx(&hash, g)
		if err != nil {
			if errors.As(err, &TxNotFoundErr{}) {
				if time.Since(timeNow) > time.Duration(waitTime)*time.Second {
					return &tx.GetTxResponse{}, TxNotFoundErr{message: "tx " + hash + " not found after " + strconv.Itoa(waitTime) + " seconds"}
				}
				time.Sleep(time.Second * time.Duration(blockTime))
				tries--
			} else {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return txRes, nil
		}
	}
	return &tx.GetTxResponse{}, RetryErr{retries: g.Retries, err: err}
}
