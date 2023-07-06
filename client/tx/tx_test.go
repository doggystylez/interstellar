package tx

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/doggystylez/interstellar/client/keys"
)

const (
	testAddress = "osmo1h0yl3rm525ay42y9a0xte4ak6tjjp3uk5s7s2y"
)

var (
	testSigningInfo = SigningInfo{
		ChainId: "osmosis-1",
		AccNum:  849140,
		SeqNum:  10,
		KeyRing: keys.KeyRing{
			KeyName:  "test",
			Backend:  "test",
			KeyBytes: []byte{221, 73, 1, 172, 141, 139, 101, 104, 210, 70, 158, 201, 52, 32, 232, 248, 75, 184, 150, 194, 105, 46, 87, 246, 149, 123, 39, 100, 111, 4, 2, 149},
		},
	}

	testTxInfo = TxInfo{
		Address:   testAddress,
		FeeAmount: 900,
		FeeDenom:  "uosmo",
		Gas:       300000,
		KeyInfo:   testSigningInfo,
	}

	testBytes = []byte{10, 137, 1, 10, 134, 1, 10, 28, 47, 99, 111, 115, 109, 111, 115, 46, 98, 97, 110, 107, 46, 118, 49, 98, 101, 116, 97, 49, 46, 77, 115, 103, 83, 101, 110, 100, 18, 102, 10, 43, 111, 115, 109, 111, 49, 104, 48, 121, 108, 51, 114, 109, 53, 50, 53, 97, 121, 52, 50, 121, 57, 97, 48, 120, 116, 101, 52, 97, 107, 54, 116, 106, 106, 112, 51, 117, 107, 53, 115, 55, 115, 50, 121, 18, 43, 111, 115, 109, 111, 49, 104, 48, 121, 108, 51, 114, 109, 53, 50, 53, 97, 121, 52, 50, 121, 57, 97, 48, 120, 116, 101, 52, 97, 107, 54, 116, 106, 106, 112, 51, 117, 107, 53, 115, 55, 115, 50, 121, 26, 10, 10, 5, 117, 111, 115, 109, 111, 18, 1, 49, 18, 102, 10, 80, 10, 70, 10, 31, 47, 99, 111, 115, 109, 111, 115, 46, 99, 114, 121, 112, 116, 111, 46, 115, 101, 99, 112, 50, 53, 54, 107, 49, 46, 80, 117, 98, 75, 101, 121, 18, 35, 10, 33, 2, 174, 244, 228, 71, 214, 193, 167, 100, 222, 57, 133, 31, 183, 41, 12, 155, 37, 179, 177, 83, 117, 47, 14, 59, 32, 159, 239, 220, 1, 234, 152, 165, 18, 4, 10, 2, 8, 1, 24, 10, 18, 18, 10, 12, 10, 5, 117, 111, 115, 109, 111, 18, 3, 57, 48, 48, 16, 224, 167, 18, 26, 64, 146, 95, 212, 76, 189, 109, 58, 92, 38, 196, 123, 217, 166, 109, 69, 109, 189, 60, 187, 42, 82, 4, 104, 49, 183, 193, 216, 168, 25, 165, 129, 16, 60, 5, 52, 209, 4, 253, 14, 252, 226, 189, 92, 153, 16, 139, 110, 246, 112, 214, 61, 59, 4, 107, 255, 132, 231, 40, 142, 209, 185, 88, 47, 191}
)

func TestSignFromPrivKey(t *testing.T) {
	txBytes, err := SignFromPrivkey([]sdk.Msg{MakeSendMsg(MsgInfo{
		From:   testAddress,
		To:     testAddress,
		Amount: 1,
		Denom:  "uosmo",
	})}, testTxInfo)
	if err != nil {
		t.Errorf("SignFromPrivkey failed with error %v", err)
	}
	if !bytes.Equal(txBytes, testBytes) {
		t.Errorf("SignFromPrivkey failed - wanted %v, got %v", testBytes, txBytes)
	}
}
