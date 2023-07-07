package query

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/doggystylez/interstellar-proto/query/account"
	"github.com/doggystylez/interstellar-proto/query/balance"
	"github.com/doggystylez/interstellar/client/grpc"
	"google.golang.org/protobuf/proto"
)

func AddressPrefix(g grpc.Client) (string, error) {
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
	return "", RetryErr{retries: g.Retries, err: err}
}

func AccountInfoFromAddress(address string, g grpc.Client) (AccountRes, error) {
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
	return AccountRes{}, RetryErr{retries: g.Retries, err: err}
}

func AllBalances(address string, g grpc.Client) (BalanceRes, error) {
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
	return BalanceRes{}, RetryErr{retries: g.Retries, err: err}

}

func BalanceByDenom(address string, denom string, g grpc.Client) (Balance, error) {
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
	return Balance{}, RetryErr{retries: g.Retries, err: err}
}
