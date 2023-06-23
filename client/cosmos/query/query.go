package query

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/doggystylez/interstellar/grpc"
	"github.com/doggystylez/interstellar/types"
	"github.com/golang/protobuf/proto"
)

func GetChainId(g types.Client) (chainId types.ChainIdRes, err error) {
	err = grpc.Open(&g)
	if err != nil {
		return
	}
	defer grpc.Close(&g)
	qyClient := tmservice.NewServiceClient(g.Conn)
	resp, err := qyClient.GetNodeInfo(g.Ctx, &tmservice.GetNodeInfoRequest{})
	if err != nil {
		return
	}
	chainId.ChainId = resp.DefaultNodeInfo.Network
	return
}

func GetAccountInfoFromAddress(config types.InterstellarConfig) (account types.AccountRes, err error) {
	err = grpc.Open(&config.Rpc)
	if err != nil {
		return
	}
	defer grpc.Close(&config.Rpc)
	qyClient := authtypes.NewQueryClient(config.Rpc.Conn)
	resp, err := qyClient.Account(config.Rpc.Ctx, &authtypes.QueryAccountRequest{Address: config.TxInfo.Address})
	if err != nil {
		return
	}
	acct := &authtypes.BaseAccount{}
	err = proto.Unmarshal(resp.Account.Value, acct)
	account.Account, account.Sequence, account.Address = acct.AccountNumber, acct.Sequence, config.TxInfo.Address
	return
}

func GetAllBalances(address string, g types.Client) (resp *banktypes.QueryAllBalancesResponse, err error) {
	err = grpc.Open(&g)
	if err != nil {
		return
	}
	defer grpc.Close(&g)
	qyClient := banktypes.NewQueryClient(g.Conn)
	resp, err = qyClient.AllBalances(g.Ctx, &banktypes.QueryAllBalancesRequest{Address: address})
	if err != nil {
		return
	}
	return
}

func GetBalanceByDenom(address string, denom string, g types.Client) (resp *banktypes.QueryBalanceResponse, err error) {
	err = grpc.Open(&g)
	if err != nil {
		return
	}
	defer grpc.Close(&g)
	qyClient := banktypes.NewQueryClient(g.Conn)
	resp, err = qyClient.Balance(g.Ctx, &banktypes.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		return
	}
	return
}

func Jsonify(input interface{}) (output string) {
	jsonData, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		panic(err)
	}
	output = string(jsonData)
	return
}
