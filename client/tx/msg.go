package tx

import (
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

func MakeSendMsg(msgInfo MsgInfo) sdk.Msg {
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	return &banktypes.MsgSend{
		FromAddress: msgInfo.From,
		ToAddress:   msgInfo.To,
		Amount:      sdk.Coins{coin},
	}

}

func MakeTransferMsg(msgInfo MsgInfo) sdk.Msg {
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	return &ibctypes.MsgTransfer{
		SourcePort:       "transfer",
		SourceChannel:    msgInfo.Channel,
		Token:            coin,
		Sender:           msgInfo.From,
		Receiver:         msgInfo.To,
		TimeoutTimestamp: uint64(time.Now().UTC().Add(+30 * time.Minute).UnixNano()),
	}

}

func MakeWasmMsg(msgInfo MsgInfo) sdk.Msg {
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	return &wasmtypes.MsgExecuteContract{
		Sender:   msgInfo.From,
		Contract: msgInfo.Contract,
		Msg:      msgInfo.ContractMsg,
		Funds:    sdk.Coins{coin},
	}
}

func msgtoMsgs(msg sdk.Msg) (msgs []sdk.Msg) {
	return []sdk.Msg{msg}
}
