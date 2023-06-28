package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	crypto "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/doggystylez/interstellar/client/keys"
)

type (
	MsgInfo struct {
		From    string
		To      string
		Amount  uint64
		Denom   string
		Channel string
		SwapMin uint64
		Routes  []SwapRoute
	}

	SwapRoute struct {
		PoolId uint64
		Denom  string
	}

	TxInfo struct {
		Address   string
		FeeAmount uint64
		FeeDenom  string
		Gas       uint64
		Memo      string
		KeyInfo   SigningInfo
	}

	SigningInfo struct {
		ChainId string
		AccNum  uint64
		SeqNum  uint64
		KeyRing keys.KeyRing
	}

	TxResponse struct {
		Code *uint32 `json:"code"`
		Hash *string `json:"hash"`
		Log  *string `json:"log"`
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
)

func NewTxConfig() TxConfig {
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	aminoCodec := codec.NewLegacyAmino()
	crypto.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
	crypto.RegisterCrypto(aminoCodec)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)
	encCfg := encodingConfig{
		InterfaceRegistry: registry,
		Codec:             cdc,
		Amino:             aminoCodec,
	}
	txBuilder := txCfg.NewTxBuilder()
	return TxConfig{
		cdc, txCfg, txBuilder, encCfg,
	}
}
