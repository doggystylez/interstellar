package address

import (
	"testing"
)

const (
	testGrpc    = "grpc.osmosis.zone:9090"
	testKeyName = "test"
	testAddress = "osmo1h0yl3rm525ay42y9a0xte4ak6tjjp3uk5s7s2y"
	testPrefix  = "osmo"
)

func TestPrefix(t *testing.T) {
	prefix, err := decodePrefix(testAddress)
	if err != nil {
		t.Errorf("DecodePrefix failed with error %v", err)
	}
	if prefix != testPrefix {
		t.Errorf("DecodePrefix failed - wanted %v, got %v", testPrefix, prefix)

	}
}
