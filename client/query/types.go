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
		Balances []Balance `json:"balances"`
	}

	Balance struct {
		Denom  string `json:"denom"`
		Amount uint64 `json:"amount"`
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
