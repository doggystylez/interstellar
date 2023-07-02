package query

import (
	"encoding/json"
	"strconv"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/doggystylez/interstellar-proto/account"
	"github.com/doggystylez/interstellar-proto/balance"
	"github.com/doggystylez/interstellar-proto/base"
	"github.com/doggystylez/interstellar-proto/wasm"
	"github.com/doggystylez/interstellar/client/grpc"
	"google.golang.org/protobuf/proto"
)

func GetChainId(g grpc.Client) (ChainIdRes, error) {
	err := g.Open()
	if err != nil {
		panic(err)

	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	res, err := client.GetNodeInfo(g.Ctx, &base.GetNodeInfoRequest{})
	if err != nil {
		return ChainIdRes{}, err
	} else {
		return ChainIdRes{ChainId: res.DefaultNodeInfo.GetNetwork()}, nil
	}
}

func GetAddressPrefix(g grpc.Client) (string, error) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := account.NewQueryClient(g.Conn)
	res, err := client.Accounts(g.Ctx, &account.QueryAccountsRequest{Pagination: &query.PageRequest{Limit: 1}})
	if err != nil {
		return "", err
	}
	acct := &account.BaseAccount{}
	err = proto.Unmarshal(res.Accounts[0].Value, acct)
	if err != nil {
		panic(err)
	}
	return decodePrefix(acct.Address), nil
}

func decodePrefix(address string) string {
	prefix, _, err := bech32.DecodeToBase256(address)
	if err != nil {
		panic(err)
	}
	return prefix
}

func GetAccountInfoFromAddress(address string, g grpc.Client) (AccountRes, error) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := account.NewQueryClient(g.Conn)
	res, err := client.Account(g.Ctx, &account.QueryAccountRequest{Address: address})
	if err != nil {
		return AccountRes{}, err
	}
	acct := &account.BaseAccount{}
	err = proto.Unmarshal(res.Account.Value, acct)
	if err != nil {
		panic(err)
	}
	return AccountRes{Address: address, Account: acct.AccountNumber, Sequence: acct.Sequence}, nil
}

func GetAllBalances(address string, g grpc.Client) (BalanceRes, error) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := balance.NewQueryClient(g.Conn)
	res, err := client.AllBalances(g.Ctx, &balance.QueryAllBalancesRequest{Address: address})
	if err != nil {
		return BalanceRes{}, err
	}
	var balances BalanceRes
	for _, coin := range res.Balances {
		amount, err := strconv.ParseUint(coin.Amount, 10, 64)
		if err != nil {
			panic(err)
		}
		balances.Balances = append(balances.Balances, Balance{Denom: coin.Denom, Amount: amount})
	}
	return balances, nil
}

func GetBalanceByDenom(address string, denom string, g grpc.Client) (Balance, error) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := balance.NewQueryClient(g.Conn)
	res, err := client.Balance(g.Ctx, &balance.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		return Balance{}, err
	}
	amount, err := strconv.ParseUint(res.Balance.Amount, 10, 64)
	if err != nil {
		panic(err)
	}
	return Balance{Denom: res.Balance.Denom, Amount: amount}, nil
}

func GetAllContractData(address string, g grpc.Client) (ContractRes, error) {
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := wasm.NewQueryClient(g.Conn)
	res, err := client.AllContractState(g.Ctx, &wasm.QueryAllContractStateRequest{Address: address, Pagination: &query.PageRequest{Limit: 2000}})
	if err != nil {
		return ContractRes{}, err
	}
	var data ContractRes
	for _, model := range res.Models {
		data.Models = append(data.Models, Model{Key: string(model.Key), Value: string(model.Value)})
	}
	return data, nil
}

func GetContractDataByQuery(address string, query interface{}, queryRes interface{}, g grpc.Client) (err error) {
	err = g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := wasm.NewQueryClient(g.Conn)
	data, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}
	res, err := client.SmartContractState(g.Ctx, &wasm.QuerySmartContractStateRequest{Address: address, QueryData: data})
	if err != nil {
		return
	} else {
		err = json.Unmarshal(res.Data, &queryRes)
		if err != nil {
			panic(err)
		}
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
