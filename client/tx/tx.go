package tx

import (
	"errors"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	crypto "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/doggystylez/interstellar/client/grpc"
	"github.com/doggystylez/interstellar/client/query"
	"github.com/doggystylez/interstellar/proto/tx"
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

func buildTx(msgs []sdk.Msg, txInfo TxInfo, txCfg *TxConfig) {
	if txInfo.FeeDenom != "" {
		feeCoin := sdk.NewCoin(txInfo.FeeDenom, sdk.NewIntFromUint64(txInfo.FeeAmount))
		txCfg.TxBuilder.SetFeeAmount(sdk.Coins{feeCoin})
	}
	txCfg.TxBuilder.SetGasLimit(txInfo.Gas)
	txCfg.TxBuilder.SetMemo(txInfo.Memo)
	err := txCfg.TxBuilder.SetMsgs(msgs...)
	if err != nil {
		panic(err)
	}
}

func broadcastTx(txBytes []byte, timeout int, g grpc.Client) (TxResponse, error) {
	var broadcastRes *tx.BroadcastTxResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	txClient := tx.NewServiceClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		broadcastRes, err = txClient.BroadcastTx(
			g.Ctx,
			&tx.BroadcastTxRequest{
				Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
				TxBytes: txBytes,
			},
		)
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			if broadcastRes.TxResponse.Code == 0 {
				if timeout == 0 {
					return TxResponse{
						Code: broadcastRes.TxResponse.Code,
						Hash: broadcastRes.TxResponse.TxHash,
						Info: "tx broadcast, not waiting for confirmation",
					}, nil
				}
				var txRes *tx.GetTxResponse
				txRes, err = query.AwaitTx(broadcastRes.TxResponse.TxHash, timeout, g)
				if err != nil {
					return TxResponse{ // nolint
						Code: broadcastRes.TxResponse.Code,
						Hash: broadcastRes.TxResponse.TxHash,
						Info: "tx broadcast but no confirmation found in " + strconv.Itoa(timeout) + " seconds",
					}, nil
				} else {
					return TxResponse{
						Code: txRes.TxResponse.Code,
						Hash: txRes.TxResponse.TxHash,
						Info: "tx confirmed",
					}, nil
				}
			} else if broadcastRes.TxResponse.Code == 19 {
				return TxResponse{
					Code: broadcastRes.TxResponse.Code,
					Hash: broadcastRes.TxResponse.TxHash,
					Info: "tx already in mempool",
				}, nil
			} else if broadcastRes.TxResponse.Code == 32 {
				return TxResponse{
					Code: broadcastRes.TxResponse.Code,
					Hash: broadcastRes.TxResponse.TxHash,
					Log:  broadcastRes.TxResponse.RawLog,
					Info: "tx sequence error",
				}, nil
			} else if broadcastRes.TxResponse.Code == 13 {
				return TxResponse{
					Code: broadcastRes.TxResponse.Code,
					Hash: broadcastRes.TxResponse.TxHash,
					Log:  broadcastRes.TxResponse.RawLog,
					Info: "insufficient fee",
				}, nil
			} else if broadcastRes.TxResponse.Code == 11 {
				return TxResponse{
					Code: broadcastRes.TxResponse.Code,
					Hash: broadcastRes.TxResponse.TxHash,
					Log:  broadcastRes.TxResponse.RawLog,
					Info: "out of gas",
				}, nil
			}
		}
	}
	return TxResponse{}, errors.New("failed sending tx after " + strconv.Itoa(g.Retries+1) + " attempts. last err: " + err.Error())
}
