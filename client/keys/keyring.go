package keys

import (
	"os"

	"github.com/99designs/keyring"
)

func Exists(keyName, path string) bool {
	if _, err := os.Stat(path + "/" + keyName); !os.IsNotExist(err) {
		return true
	}
	return false
}

func Save(keyName string, path string, keyBytes []byte, password string) (err error) {
	var passFunc keyring.PromptFunc
	if password == "" {
		passFunc = keyring.TerminalPrompt
	} else {
		passFunc = keyring.FixedStringPrompt(password)
	}
	ring, err := keyring.Open(keyring.Config{
		ServiceName:      "interstellar",
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		FileDir:          path,
		FilePasswordFunc: passFunc,
	})
	if err != nil {
		return
	}
	err = ring.Set(keyring.Item{
		Key:  keyName,
		Data: keyBytes,
	})
	return
}

func Load(keyName, path, password string) (keyBytes []byte, err error) {
	var passFunc keyring.PromptFunc
	if password == "" {
		passFunc = keyring.TerminalPrompt
	} else {
		passFunc = keyring.FixedStringPrompt(password)
	}
	ring, err := keyring.Open(keyring.Config{
		ServiceName:      "interstellar",
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		FileDir:          path,
		FilePasswordFunc: passFunc,
	})
	if err != nil {
		return
	}
	key, err := ring.Get(keyName)
	if err != nil {
		return
	}
	keyBytes = key.Data
	return
}
