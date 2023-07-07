package query

import (
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
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.Accounts(g.Ctx, &account.QueryAccountsRequest{Pagination: &query.PageRequest{Limit: 1}})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
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
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.Account(g.Ctx, &account.QueryAccountRequest{Address: address})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
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
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.AllBalances(g.Ctx, &balance.QueryAllBalancesRequest{Address: address})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			var balances BalanceRes
			for _, coin := range res.Balances {
				balances.Balances = append(balances.Balances, Token{Denom: coin.Denom, Amount: coin.Amount})
			}
			return balances, nil
		}
	}
	return BalanceRes{}, RetryErr{retries: g.Retries, err: err}

}

func BalanceByDenom(address string, denom string, g grpc.Client) (Token, error) {
	var res *balance.QueryBalanceResponse
	err := g.Open()
	if err != nil {
		panic(err)
	}
	defer g.Close()
	client := balance.NewQueryClient(g.Conn)
	tries, maxTries := 0, g.Retries+1
	for tries < maxTries {
		tries++
		res, err = client.Balance(g.Ctx, &balance.QueryBalanceRequest{Address: address, Denom: denom})
		if err != nil {
			if tries < maxTries {
				time.Sleep(time.Duration(g.Interval) * time.Second)
			}
		} else {
			return Token{Denom: res.Balance.Denom, Amount: res.Balance.Amount}, nil
		}
	}
	return Token{}, RetryErr{retries: g.Retries, err: err}
}
