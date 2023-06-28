package tx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/doggystylez/interstellar/client/grpc"
)

func AssembleAndBroadcast(msgInfo MsgInfo, txInfo TxInfo, rpc grpc.Client, maker MsgMaker) (res TxResponse, err error) {
	txBytes, err := SignFromPrivkey(msgtoMsgs(maker(msgInfo)), txInfo)
	if err != nil {
		return
	}
	res, err = broadcastTx(txBytes, rpc)
	return
}

func buildTx(msgs []sdk.Msg, txInfo TxInfo) (txConfig TxConfig, err error) {
	txConfig = NewTxConfig()
	if txInfo.FeeDenom != "" {
		feeCoin := sdk.NewCoin(txInfo.FeeDenom, sdk.NewIntFromUint64(txInfo.FeeAmount))
		txConfig.TxBuilder.SetFeeAmount(sdk.Coins{feeCoin})
	}
	txConfig.TxBuilder.SetGasLimit(txInfo.Gas)
	txConfig.TxBuilder.SetMemo(txInfo.Memo)
	err = txConfig.TxBuilder.SetMsgs(msgs...)
	return
}

func broadcastTx(txBytes []byte, g grpc.Client) (resp TxResponse, err error) {
	err = g.Open()
	if err != nil {
		return
	}
	defer g.Close()
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
	resp = TxResponse{
		Code: &grpcRes.TxResponse.Code,
		Hash: &grpcRes.TxResponse.TxHash,
		Log:  &grpcRes.TxResponse.RawLog,
	}
	return
}
