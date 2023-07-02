package query

import (
	"testing"

	"github.com/doggystylez/interstellar/client/grpc"
)

const (
	testGrpc    = "grpc.osmosis.zone:9090"
	testAddress = "osmo1h0yl3rm525ay42y9a0xte4ak6tjjp3uk5s7s2y"
	testPrefix  = "osmo"
	testchainId = "osmosis-1"
	testAccount = 849140
)

var testClient = grpc.Client{
	Endpoint: testGrpc,
	Timeout:  7,
}

func TestChainID(t *testing.T) {
	err := testClient.Open()
	if err != nil {
		panic(err)
	}
	defer testClient.Close()
	chainId, _ := GetChainId(testClient)
	if chainId.ChainId != testchainId {
		t.Errorf("GetChainId failed - wanted %v, got %v", testchainId, chainId.ChainId)

	}
}

func TestAccountInfo(t *testing.T) {
	err := testClient.Open()
	if err != nil {
		panic(err)
	}
	defer testClient.Close()
	account, _ := GetAccountInfoFromAddress(testAddress, testClient)
	if account.Account != testAccount {
		t.Errorf("getaccount failed - wanted %v, got %v", testAccount, account.Account)

	}
}

func TestGetPrefix(t *testing.T) {
	err := testClient.Open()
	if err != nil {
		panic(err)
	}
	defer testClient.Close()
	prefix, _ := GetAddressPrefix(testClient)
	if prefix != testPrefix {
		t.Errorf("getprefix failed - wanted %v, got %v", testPrefix, prefix)

	}
}
