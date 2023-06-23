package address

import (
	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/doggystylez/interstellar/grpc"
	"github.com/doggystylez/interstellar/types"
	"github.com/golang/protobuf/proto"
)

func GetAddressPrefix(g types.Client) (prefix string, err error) {
	err = grpc.Open(&g)
	if err != nil {
		panic(err)
	}
	defer grpc.Close(&g)
	qyClient := authtypes.NewQueryClient(g.Conn)
	qy := authtypes.QueryAccountsRequest{}
	qy.Pagination = &query.PageRequest{Limit: 1}
	resp, err := qyClient.Accounts(g.Ctx, &qy)
	if err != nil {
		return
	}
	acct := &authtypes.BaseAccount{}
	err = proto.Unmarshal(resp.Accounts[0].Value, acct)
	if err != nil {
		return
	}
	prefix, err = decodePrefix(acct.Address)
	return
}

func decodePrefix(address string) (prefix string, err error) {
	prefix, _, err = bech32.DecodeToBase256(address)
	return
}
