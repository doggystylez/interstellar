package keys

type (
	KeyRing struct {
		KeyName  string
		Backend  string
		KeyBytes []byte
	}

	Addrbook struct {
		InternalAccounts  []Entry `json:"internal,omitempty"`
		WithdrawAddresses []Entry `json:"external,omitempty"`
	}

	Entry struct {
		Name     string             `json:"name,omitempty"`
		Accounts map[string]Account `json:"accounts,omitempty"`
	}

	Account struct {
		Address string `json:"address,omitempty"`
		AccNum  uint64 `json:"acc_num,omitempty"`
		SeqNum  uint64 `json:"seq_num,omitempty"`
	}
)
