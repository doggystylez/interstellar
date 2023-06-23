package tx

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/doggystylez/interstellar/grpc"
	"github.com/doggystylez/interstellar/keys/keyring"
	"github.com/doggystylez/interstellar/types"
)

func AssembleAndBroadcast(msgInfo types.MsgInfo, config types.InterstellarConfig, maker types.MsgMaker) (res types.TxResponse, err error) {
	if config.TxInfo.KeyInfo.KeyRing.HexPriv == "" {
		var keyBytes []byte
		keyBytes, err = keyring.Load(config)
		if err != nil {
			return
		}
		config.TxInfo.KeyInfo.KeyRing.HexPriv = hex.EncodeToString(keyBytes)
	}
	txBytes, err := SignFromPrivkey(msgtoMsgs(maker(msgInfo)), config)
	if err != nil {
		return
	}
	res, err = broadcastTx(txBytes, config.Rpc)
	return
}

func buildTx(msgs []sdk.Msg, txInfo types.TxInfo) (txConfig types.TxConfig, err error) {
	txConfig = types.NewTxConfig()
	if txInfo.FeeDenom != "" {
		feeCoin := sdk.NewCoin(txInfo.FeeDenom, sdk.NewIntFromUint64(txInfo.FeeAmount))
		txConfig.TxBuilder.SetFeeAmount(sdk.Coins{feeCoin})
	}
	txConfig.TxBuilder.SetGasLimit(txInfo.Gas)
	txConfig.TxBuilder.SetMemo(txInfo.Memo)
	err = txConfig.TxBuilder.SetMsgs(msgs...)
	return
}

func broadcastTx(txBytes []byte, g types.Client) (resp types.TxResponse, err error) {
	err = grpc.Open(&g)
	if err != nil {
		return
	}
	defer grpc.Close(&g)
	txClient := tx.NewServiceClient(g.Conn)
	grpcRes, err := txClient.BroadcastTx(
		g.Ctx,
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return
	}
	resp = types.TxResponse{
		Code: &grpcRes.TxResponse.Code,
		Hash: &grpcRes.TxResponse.TxHash,
		Log:  &grpcRes.TxResponse.RawLog,
	}
	return
}
