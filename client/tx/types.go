package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/doggystylez/interstellar/client/keys"
)

type (
	MsgInfo struct {
		From        string
		To          string
		Amount      string
		Denom       string
		Channel     string
		Contract    string
		ContractMsg []byte
		Maker       MsgMaker
	}

	TxInfo struct {
		Address        string
		FeeAmount      uint64
		FeeDenom       string
		Gas            uint64
		Memo           string
		ConfirmTimeout int
		KeyInfo        SigningInfo
	}

	SigningInfo struct {
		ChainId string
		AccNum  uint64
		SeqNum  uint64
		KeyRing keys.KeyRing
	}

	TxResponse struct {
		Code uint32 `json:"code"`
		Hash string `json:"hash"`
		Log  string `json:"log"`
	}

	TxConfig struct {
		Codec          codec.Codec
		TxConfig       client.TxConfig
		TxBuilder      client.TxBuilder
		EncodingConfig encodingConfig
	}

	encodingConfig struct {
		InterfaceRegistry codectypes.InterfaceRegistry
		Codec             codec.Codec
		Amino             *codec.LegacyAmino
	}

	MsgMaker func(MsgInfo) sdk.Msg

	WasmSwap struct {
		Swap `json:"swap"`
	}

	Swap struct {
		InputCoin   `json:"input_coin"`
		OutputDenom string `json:"output_denom"`
		Slippage    `json:"slippage"`
	}

	InputCoin struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	Slippage struct {
		Twap `json:"twap"`
	}

	Twap struct {
		SlippagePercentage string `json:"slippage_percentage"`
		WindowSeconds      int    `json:"window_seconds"`
	}
)
