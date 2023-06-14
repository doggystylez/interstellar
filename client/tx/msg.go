package tx

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/doggystylez/interstellar/types"
)

func MakeSendMsg(msgInfo types.MsgInfo) (msg sdk.Msg) {
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	msg = &banktypes.MsgSend{
		FromAddress: msgInfo.From,
		ToAddress:   msgInfo.To,
		Amount:      sdk.Coins{coin},
	}
	return
}

func MakeTransferMsg(msgInfo types.MsgInfo) (msg sdk.Msg) {
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	msg = &ibctypes.MsgTransfer{
		SourcePort:       "transfer",
		SourceChannel:    msgInfo.Channel,
		Token:            coin,
		Sender:           msgInfo.From,
		Receiver:         msgInfo.To,
		TimeoutTimestamp: uint64(time.Now().UTC().Add(+30 * time.Minute).UnixNano()),
	}
	return
}

func msgtoMsgs(msg sdk.Msg) (msgs []sdk.Msg) {
	return []sdk.Msg{msg}
}
