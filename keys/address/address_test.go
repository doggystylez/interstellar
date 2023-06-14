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

// var (
// 	testSigningInfo = types.SigningInfo{
// 		ChainId: "osmosis-1",
// 		AccNum:  849140,
// 		SeqNum:  10,
// 		KeyRing: types.KeyRing{
// 			KeyName: "test",
// 			Backend: "test",
// 			//		KeyPath: "/home/debian/Code/git/interstellar",
// 			HexPriv: "dd4901ac8d8b6568d2469ec93420e8f84bb896c2692e57f6957b27646f040295",
// 		},
// 	}
// )

func TestPrefix(t *testing.T) {
	prefix, err := decodePrefix(testAddress)
	if err != nil {
		t.Errorf("DecodePrefix failed with error %v", err)
	}
	if prefix != testPrefix {
		t.Errorf("DecodePrefix failed - wanted %v, got %v", testPrefix, prefix)

	}
}

// func TestPrivDerivations(t *testing.T) {
// 	add, err := encodeAddressFromPriv(testPrefix, testSigningInfo.KeyRing.HexPriv)
// 	if err != nil {
// 		t.Errorf("AddressfromPrivKey failed with error %v", err)
// 	}
// 	if add != testAddress {
// 		t.Errorf("AddressfromPrivKey failed - wanted %v, got %v", testAddress, add)

// 	}
// }
