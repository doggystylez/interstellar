package tx

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	pooltypes "github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"
)

func MakeSendMsg(msgInfo MsgInfo) (msg sdk.Msg) {
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	msg = &banktypes.MsgSend{
		FromAddress: msgInfo.From,
		ToAddress:   msgInfo.To,
		Amount:      sdk.Coins{coin},
	}
	return
}

func MakeTransferMsg(msgInfo MsgInfo) (msg sdk.Msg) {
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

func MakeSwapMsg(msgInfo MsgInfo) (msg sdk.Msg) {
	var routes []pooltypes.SwapAmountInRoute
	coin := sdk.NewCoin(msgInfo.Denom, sdk.NewIntFromUint64(msgInfo.Amount))
	for i, route := range msgInfo.Routes {
		routes[i] = pooltypes.SwapAmountInRoute{PoolId: route.PoolId, TokenOutDenom: route.Denom}
	}
	msg = &pooltypes.MsgSwapExactAmountIn{
		Sender:            msgInfo.From,
		Routes:            routes,
		TokenIn:           coin,
		TokenOutMinAmount: sdk.NewIntFromUint64(msgInfo.SwapMin),
	}
	return
}

func msgtoMsgs(msg sdk.Msg) (msgs []sdk.Msg) {
	return []sdk.Msg{msg}
}
