package keys

type (
	KeyRing struct {
		KeyName string
		Backend string
		//	HexPriv  string
		Mnemonic string
		KeyBytes []byte
	}
)
