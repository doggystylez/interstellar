package query

import (
	"testing"

	"github.com/doggystylez/interstellar/grpc"
	"github.com/doggystylez/interstellar/keys/address"
	"github.com/doggystylez/interstellar/types"
)

const (
	testGrpc    = "grpc.osmosis.zone:9090"
	testAddress = "osmo1h0yl3rm525ay42y9a0xte4ak6tjjp3uk5s7s2y"
	testPrefix  = "osmo"
	testchainId = "osmosis-1"
	testAccount = 849140
)

var testConfig = types.InterstellarConfig{
	Rpc:    testClient,
	TxInfo: types.TxInfo{Address: testAddress},
}

var testClient = types.Client{
	Endpoint: testGrpc,
	Timeout:  7,
}

func TestChainID(t *testing.T) {
	err := grpc.Open(&testClient)
	if err != nil {
		panic(err)
	}
	defer grpc.Close(&testClient)
	chainId, err := GetChainId(testClient)
	if err != nil {
		panic(err)
	}
	if chainId.ChainId != testchainId {
		t.Errorf("GetChainId failed - wanted %v, got %v", testchainId, chainId.ChainId)

	}
}

func TestAccountInfo(t *testing.T) {
	err := grpc.Open(&testClient)
	if err != nil {
		panic(err)
	}
	defer grpc.Close(&testClient)
	account, err := GetAccountInfoFromAddress(testConfig)
	if err != nil {
		panic(err)
	}
	if account.Account != testAccount {
		t.Errorf("getaccount failed - wanted %v, got %v", testAccount, account.Account)

	}
}

func TestGetPrefix(t *testing.T) {
	err := grpc.Open(&testClient)
	if err != nil {
		panic(err)
	}
	defer grpc.Close(&testClient)
	prefix, err := address.GetAddressPrefix(testClient)
	if err != nil {
		panic(err)
	}
	if prefix != testPrefix {
		t.Errorf("getprefix failed - wanted %v, got %v", testPrefix, prefix)

	}
}
