package tx

import (
	"errors"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	crypto "github.com/cosmos/cosmos-sdk/crypto/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/doggystylez/interstellar-proto/tx"
	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/doggystylez/interstellar/client/query"
)

var txConfig TxConfig

func init() {
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	aminoCodec := codec.NewLegacyAmino()
	crypto.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
	crypto.RegisterCrypto(aminoCodec)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)
	txConfig = TxConfig{
		cdc, txCfg, txCfg.NewTxBuilder(), encodingConfig{
			InterfaceRegistry: registry,
			Codec:             cdc,
			Amino:             aminoCodec,
		},
	}
}

func AssembleAndBroadcast(msgs []MsgInfo, txInfo TxInfo, rpc grpc.Client) (TxResponse, error) {
	txBytes, err := SignFromPrivkey(makeMsgs(msgs), txInfo)
	if err != nil {
		return TxResponse{}, err
	}
	return broadcastTx(txBytes, txInfo.ConfirmTimeout, rpc)
}

func broadcastTx(txBytes []byte, timeout int, g grpc.Client) (TxResponse, error) {
	var broadcastRes *tx.BroadcastTxResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	txClient := tx.NewServiceClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		broadcastRes, err = txClient.BroadcastTx(
			g.Ctx,
			&tx.BroadcastTxRequest{
				Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
				TxBytes: txBytes,
			},
		)
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			if broadcastRes.TxResponse.Code == 0 {
				if timeout == 0 {
					return TxResponse{
						Code: broadcastRes.TxResponse.Code,
						Hash: broadcastRes.TxResponse.TxHash,
						Log:  "tx broadcast, not waiting for confirmation",
					}, nil
				}
				var txRes *tx.GetTxResponse
				txRes, err = query.AwaitTx(broadcastRes.TxResponse.TxHash, timeout, g)
				if err != nil {
					return TxResponse{ // nolint
						Code: broadcastRes.TxResponse.Code,
						Hash: broadcastRes.TxResponse.TxHash,
						Log:  "tx broadcast but no confirmation found in " + strconv.Itoa(timeout) + " seconds",
					}, nil
				} else {
					return TxResponse{
						Code: txRes.TxResponse.Code,
						Hash: txRes.TxResponse.TxHash,
						Log:  "tx confirmed",
					}, nil
				}
			} else if broadcastRes.TxResponse.Code == 19 {
				return TxResponse{
					Code: broadcastRes.TxResponse.Code,
					Hash: broadcastRes.TxResponse.TxHash,
					Log:  "tx already in mempool",
				}, nil
			} else if broadcastRes.TxResponse.Code == 32 {
				return TxResponse{
					Code: broadcastRes.TxResponse.Code,
					Hash: broadcastRes.TxResponse.TxHash,
					Log:  "tx sequence error",
				}, nil
			}
		}
	}
	return TxResponse{}, errors.New("failed sending tx after " + strconv.Itoa(g.Retries+1) + " attempts. last err: " + err.Error())
}
