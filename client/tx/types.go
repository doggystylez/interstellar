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
		Code uint32 `json:"code,omitempty"`
		Hash string `json:"hash,omitempty"`
		Log  string `json:"log,omitempty"`
		Info string `json:"info,omitempty"`
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

	Swap struct {
		InputCoin         *Coin  `json:"input_coin,omitempty"`
		OutputDenom       string `json:"output_denom,omitempty"`
		*OnFailedDelivery `json:"on_failed_delivery,omitempty"`
		Receiver          string `json:"receiver,omitempty"`
		Slippage          `json:"slippage,omitempty"`
	}

	Coin struct {
		Denom  string `json:"denom,omitempty"`
		Amount string `json:"amount,omitempty"`
	}
	Slippage struct {
		Twap `json:"twap,omitempty"`
	}

	Twap struct {
		SlippagePercentage string `json:"slippage_percentage"`
		WindowSeconds      int    `json:"window_seconds"`
	}

	ContractSwap struct {
		Swap `json:"swap,omitempty"`
	}

	ContractIbcSwap struct {
		OsmosisSwap Swap `json:"osmosis_swap,omitempty"`
	}

	OnFailedDelivery struct {
		Addr string `json:"local_recovery_addr,omitempty"`
	}
)
