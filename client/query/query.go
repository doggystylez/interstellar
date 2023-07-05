package query

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/doggystylez/interstellar-proto/account"
	"github.com/doggystylez/interstellar-proto/balance"
	"github.com/doggystylez/interstellar-proto/base"
	"github.com/doggystylez/interstellar-proto/wasm"
	"github.com/doggystylez/interstellar/client/grpc"
	"google.golang.org/protobuf/proto"
)

func (e retryErr) Error() string {
	return "failed to query after " + strconv.Itoa(e.retries+1) + " attempts. last error: " + e.err.Error()
}

func decodePrefix(address string) string {
	prefix, _, err := bech32.DecodeToBase256(address)
	if err != nil {
		panic(err)
	}
	return prefix
}

func GetChainId(g grpc.Client) (ChainIdRes, error) {
	var res *base.GetNodeInfoResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := base.NewServiceClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.GetNodeInfo(g.Ctx, &base.GetNodeInfoRequest{})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			return ChainIdRes{ChainId: res.DefaultNodeInfo.GetNetwork()}, nil
		}
	}
	return ChainIdRes{}, retryErr{retries: g.Retries, err: err}
}

func GetAddressPrefix(g grpc.Client) (string, error) {
	var res *account.QueryAccountsResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := account.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.Accounts(g.Ctx, &account.QueryAccountsRequest{Pagination: &query.PageRequest{Limit: 1}})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			acct := &account.BaseAccount{}
			err = proto.Unmarshal(res.Accounts[0].Value, acct)
			if err != nil {
				panic(err)
			}
			return decodePrefix(acct.Address), nil
		}
	}
	return "", retryErr{retries: g.Retries, err: err}
}

func GetAccountInfoFromAddress(address string, g grpc.Client) (AccountRes, error) {
	var res *account.QueryAccountResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := account.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.Account(g.Ctx, &account.QueryAccountRequest{Address: address})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			acct := &account.BaseAccount{}
			err = proto.Unmarshal(res.Account.Value, acct)
			if err != nil {
				panic(err)
			}
			return AccountRes{Address: address, Account: acct.AccountNumber, Sequence: acct.Sequence}, nil
		}
	}
	return AccountRes{}, retryErr{retries: g.Retries, err: err}
}

func GetAllBalances(address string, g grpc.Client) (BalanceRes, error) {
	var res *balance.QueryAllBalancesResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := balance.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.AllBalances(g.Ctx, &balance.QueryAllBalancesRequest{Address: address})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			var (
				balances BalanceRes
				amount   uint64
			)
			for _, coin := range res.Balances {
				amount, err = strconv.ParseUint(coin.Amount, 10, 64)
				if err != nil {
					panic(err)
				}
				balances.Balances = append(balances.Balances, Balance{Denom: coin.Denom, Amount: amount})
			}
			return balances, nil
		}
	}
	return BalanceRes{}, retryErr{retries: g.Retries, err: err}

}

func GetBalanceByDenom(address string, denom string, g grpc.Client) (Balance, error) {
	var res *balance.QueryBalanceResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := balance.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.Balance(g.Ctx, &balance.QueryBalanceRequest{Address: address, Denom: denom})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			var amount uint64
			amount, err = strconv.ParseUint(res.Balance.Amount, 10, 64)
			if err != nil {
				panic(err)
			}
			return Balance{Denom: res.Balance.Denom, Amount: amount}, nil
		}
	}
	return Balance{}, retryErr{retries: g.Retries, err: err}
}

func GetAllContractData(address string, g grpc.Client) (ContractRes, error) {
	var res *wasm.QueryAllContractStateResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := wasm.NewQueryClient(g.Conn)
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.AllContractState(g.Ctx, &wasm.QueryAllContractStateRequest{Address: address, Pagination: &query.PageRequest{Limit: 2000}})
		if err != nil {
			time.Sleep(time.Duration(g.Interval) * time.Second)
		} else {
			var data ContractRes
			for _, model := range res.Models {
				data.Models = append(data.Models, Model{Key: model.Key, Value: model.Value})
			}
			return data, nil
		}
	}
	return ContractRes{}, retryErr{retries: g.Retries, err: err}
}

func GetContractDataByQuery(address string, query interface{}, queryRes interface{}, g grpc.Client) error {
	var res *wasm.QuerySmartContractStateResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := wasm.NewQueryClient(g.Conn)
	data, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}
	for tries := -1; tries < g.Retries; tries++ {
		res, err = client.SmartContractState(g.Ctx, &wasm.QuerySmartContractStateRequest{Address: address, QueryData: data})
		if err != nil {
			if strings.Contains(err.Error(), "query wasm contract failed") {
				return err
			} else {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			err = json.Unmarshal(res.Data, &queryRes)
			if err != nil {
				panic(err)
			}
			return nil
		}
	}
	return retryErr{retries: g.Retries, err: err}
}

func Jsonify(input interface{}) (output string) {
	jsonData, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		panic(err)
	}
	output = string(jsonData)
	return
}
