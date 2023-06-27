package query

import (
	"encoding/json"
	"strconv"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/doggystylez/interstellar-proto/account"
	"github.com/doggystylez/interstellar-proto/balance"
	"github.com/doggystylez/interstellar-proto/base"
	"github.com/doggystylez/interstellar/client/grpc"
	"google.golang.org/protobuf/proto"
)

func GetChainId(g grpc.Client) (chain ChainIdRes) {
	err := g.Open()
	if err != nil {
		panic(err)

	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	res, err := client.GetNodeInfo(g.Ctx, &base.GetNodeInfoRequest{})
	if err != nil {
		panic(err)
	}
	chain.ChainId = res.DefaultNodeInfo.Network
	return
}

func GetAddressPrefix(g grpc.Client) (prefix string) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := account.NewQueryClient(g.Conn)
	res, err := client.Accounts(g.Ctx, &account.QueryAccountsRequest{Pagination: &query.PageRequest{Limit: 1}})
	if err != nil {
		panic(err)
	}
	acct := &account.BaseAccount{}
	err = proto.Unmarshal(res.Accounts[0].Value, acct)
	if err != nil {
		panic(err)
	}
	prefix, err = decodePrefix(acct.Address)
	if err != nil {
		panic(err)
	}
	return
}

func decodePrefix(address string) (prefix string, err error) {
	prefix, _, err = bech32.DecodeToBase256(address)
	return
}

func GetAccountInfoFromAddress(address string, g grpc.Client) (a AccountRes) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := account.NewQueryClient(g.Conn)
	res, err := client.Account(g.Ctx, &account.QueryAccountRequest{Address: address})
	if err != nil {
		panic(err)
	}
	acct := &account.BaseAccount{}
	err = proto.Unmarshal(res.Account.Value, acct)
	if err != nil {
		panic(err)
	}
	a.Account, a.Sequence, a.Address = acct.AccountNumber, acct.Sequence, address
	return
}

func GetAllBalances(address string, g grpc.Client) (b BalanceRes) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()

	client := balance.NewQueryClient(g.Conn)
	res, err := client.AllBalances(g.Ctx, &balance.QueryAllBalancesRequest{Address: address})
	if err != nil {
		panic(err)
	}
	for _, coin := range res.Balances {
		amount, err := strconv.ParseUint(coin.Amount, 10, 64)
		if err != nil {
			panic(err)
		}
		b.Balances = append(b.Balances, Balance{Denom: coin.Denom, Amount: amount})
	}
	return
}

func GetBalanceByDenom(address string, denom string, g grpc.Client) (b Balance) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := balance.NewQueryClient(g.Conn)
	res, err := client.Balance(g.Ctx, &balance.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		panic(err)
	}
	amount, err := strconv.ParseUint(res.Balance.Amount, 10, 64)
	if err != nil {
		panic(err)
	}
	b.Denom, b.Amount = res.Balance.Denom, amount
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
