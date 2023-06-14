package keyring

import (
	"fmt"

	"github.com/99designs/keyring"
	"github.com/doggystylez/interstellar/types"
)

func Save(config types.InterstellarConfig, keyBytes []byte) (err error) {
	ring, err := keyring.Open(keyring.Config{
		ServiceName:      "interstellar",
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		FileDir:          config.Path,
		FilePasswordFunc: keyring.TerminalPrompt,
	})
	if err != nil {
		return
	}
	err = ring.Set(keyring.Item{
		Key:  config.TxInfo.KeyInfo.KeyRing.KeyName,
		Data: keyBytes,
	})
	if err == nil {
		fmt.Println("saved key to keyring")
	}
	return
}

func Load(config types.InterstellarConfig) (keyBytes []byte, err error) {
	ring, err := keyring.Open(keyring.Config{
		ServiceName:      "interstellar",
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		FileDir:          config.Path,
		FilePasswordFunc: keyring.TerminalPrompt,
	})
	if err != nil {
		return
	}
	key, err := ring.Get(config.TxInfo.KeyInfo.KeyRing.KeyName)
	if err != nil {
		return
	}
	keyBytes = key.Data
	return
}
