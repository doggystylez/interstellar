package types

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
)
