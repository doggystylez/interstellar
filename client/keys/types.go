package keys

type (
	KeyRing struct {
		KeyName  string
		Backend  string
		KeyBytes []byte
	}

	Addrbook struct {
		InternalAddresses []Entry `json:"internal,omitempty"`
		WithdrawAddresses []Entry `json:"external,omitempty"`
	}

	Entry struct {
		Name      string            `json:"name,omitempty"`
		Addresses map[string]string `json:"addresses,omitempty"`
	}
)
