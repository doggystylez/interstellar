package tx

import (
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

func MakeSendMsg(msgInfo MsgInfo) sdk.Msg {
	amount, ok := sdk.NewIntFromString(msgInfo.Amount)
	if !ok {
		panic("conversion from " + msgInfo.Amount + " to big.Int failed")
	}
	coin := sdk.NewCoin(msgInfo.Denom, amount)
	return &banktypes.MsgSend{
		FromAddress: msgInfo.From,
		ToAddress:   msgInfo.To,
		Amount:      sdk.Coins{coin},
	}
}

func MakeTransferMsg(msgInfo MsgInfo) sdk.Msg {
	amount, ok := sdk.NewIntFromString(msgInfo.Amount)
	if !ok {
		panic("conversion from " + msgInfo.Amount + " to big.Int failed")
	}
	coin := sdk.NewCoin(msgInfo.Denom, amount)
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
	msg := &wasmtypes.MsgExecuteContract{
		Sender:   msgInfo.From,
		Contract: msgInfo.Contract,
		Msg:      msgInfo.ContractMsg,
	}
	if msgInfo.Amount != "" {
		amount, ok := sdk.NewIntFromString(msgInfo.Amount)
		if !ok {
			panic("conversion from " + msgInfo.Amount + " to big.Int failed")
		}
		msg.Funds = sdk.Coins{sdk.NewCoin(msgInfo.Denom, amount)}
	}
	return msg
}

func makeMsgs(msgs []MsgInfo) (sdkMsgs []sdk.Msg) {
	for _, msg := range msgs {
		sdkMsgs = append(sdkMsgs, msg.Maker(msg))
	}
	return
}
