package query

type (
	ChainIdRes struct {
		ChainId string `json:"chain_id"`
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

	PoolsRes struct {
		Balancer   []BalancerPool   `json:"balancer"`
		Stableswap []StableswapPool `json:"stableswap"`
	}

	BalancerPool struct {
		Id         uint64      `json:"id,omitempty"`
		PoolAssets []PoolAsset `json:"pool_assets,omitempty"`
	}

	StableswapPool struct {
		Id             uint64   `json:"id"`
		PoolLiquidity  []Token  `json:"pool_liquidity"`
		ScalingFactors []uint64 `json:"scaling_factors"`
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
