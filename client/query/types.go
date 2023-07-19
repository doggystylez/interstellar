package query

import "github.com/cometbft/cometbft/proto/tendermint/types"

type (
	ChainIdRes struct {
		ChainId string `json:"chain_id"`
	}

	BlockRes struct {
		Block types.Block `json:"block"`
	}

	AddressRes struct {
		Address string `json:"address"`
	}

	AccountRes struct {
		Address  string `json:"address"`
		Account  uint64 `json:"account"`
		Sequence uint64 `json:"sequence"`
	}

	BalanceRes struct {
		Balances []Token `json:"balances"`
	}

	Token struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	ContractRes struct {
		Models []Model `json:"models"`
	}

	Model struct {
		Key   []byte `json:"key"`
		Value []byte `json:"value"`
	}

	RetryErr struct {
		retries int
		err     error
	}

	TxNotFoundErr struct {
		message string
	}
)
