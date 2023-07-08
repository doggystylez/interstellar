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

	SpotPriceRes struct {
		Base  string `json:"base"`
		Quote string `json:"quote"`
		Price string `json:"price"`
	}

	SwapEstimate struct {
		TokenIn  Token `json:"token_in"`
		TokenOut Token `json:"token_out"`
	}

	SwapRoute struct {
		PoolId   uint64 `json:"pool_id"`
		DenomOut string `json:"token_out"`
	}

	PoolsRes []Pool

	Pool struct {
		Id             uint64      `json:"id,omitempty"`
		Type           string      `json:"type,omitempty"`
		PoolAssets     []PoolAsset `json:"pool_assets,omitempty"`
		PoolLiquidity  []Token     `json:"pool_liquidity,omitempty"`
		ScalingFactors []uint64    `json:"scaling_factors,omitempty"`
	}

	PoolAsset struct {
		Token  `json:"token,omitempty"`
		Weight string `json:"weight,omitempty"`
	}

	RetryErr struct {
		retries int
		err     error
	}

	TxNotFoundErr struct {
		message string
	}
)
