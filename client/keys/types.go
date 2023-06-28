package keys

type (
	KeyRing struct {
		KeyName  string
		Backend  string
		Mnemonic string
		KeyBytes []byte
	}
)
