package query

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/doggystylez/interstellar/proto/tx"
)

func Tx(hash *string, g grpc.Client) (*tx.GetTxResponse, error) {
	var res *tx.GetTxResponse
	*hash = strings.ToUpper(*hash)
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := tx.NewServiceClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.GetTx(g.Ctx, &tx.GetTxRequest{Hash: *hash})
		if err != nil {
			if strings.Contains(err.Error(), *hash) && strings.Contains(err.Error(), "not found") {
				return &tx.GetTxResponse{}, TxNotFoundErr{message: "tx " + *hash + " not found"}
			}
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
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
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		txRes, err = Tx(&hash, g)
		if err != nil {
			if errors.As(err, &TxNotFoundErr{}) {
				if time.Since(timeNow) > time.Duration(waitTime)*time.Second {
					return &tx.GetTxResponse{}, TxNotFoundErr{message: "tx " + hash + " not found after " + strconv.Itoa(waitTime) + " seconds"}
				}
				time.Sleep(time.Second * time.Duration(blockTime))
				tries--
			} else {
				if tries < maxTries {
					time.Sleep(time.Duration(g.Interval) * time.Second)
				}
			}
		} else {
			return txRes, nil
		}
	}
	return &tx.GetTxResponse{}, RetryErr{retries: g.Retries, err: err}
}
