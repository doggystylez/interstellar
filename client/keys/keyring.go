package keys

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/keyring"
	"golang.org/x/term"
)

func passPrompt(keyName string) keyring.PromptFunc {
	fmt.Println("enter password for key", keyName+":") // nolint
	return func(_ string) (string, error) {
		pass, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		return strings.TrimSuffix(string(pass), string('\n')), nil
	}
}

func ensureKeyring(path, backend string) {
	var keyDir string
	if backend == "file" {
		keyDir = filepath.Join(path, "keyring-file")
	} else if backend == "test" {
		keyDir = filepath.Join(path, "keyring-test")
	}
	if _, err := os.Stat(keyDir); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			panic(err)
		}
	}
}

func Exists(keyName, path, backend string) bool {
	var keyFile string
	if backend == "file" {
		keyFile = filepath.Join(path, "keyring-file", keyName)
	} else if backend == "test" {
		keyFile = filepath.Join(path, "keyring-test", keyName)
	}
	if _, err := os.Stat(keyFile); !os.IsNotExist(err) {
		return true
	}
	return false
}

func Save(keyBytes []byte, keyName string, path string, backend string, password string) (err error) {
	var (
		keyDir   string
		passFunc keyring.PromptFunc
	)
	ensureKeyring(path, backend)
	if backend == "file" {
		keyDir = filepath.Join(path, "keyring-file")
	} else if backend == "test" {
		keyDir = filepath.Join(path, "keyring-test")
		password = "interstellar"
	}
	if password == "" {
		passFunc = passPrompt(keyName)
	} else {
		passFunc = keyring.FixedStringPrompt(password)
	}
	ring, err := keyring.Open(keyring.Config{
		ServiceName:      "interstellar",
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		FileDir:          keyDir,
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

func Load(keyName, path, backend, password string) (keyBytes []byte, err error) {
	var (
		keyDir   string
		passFunc keyring.PromptFunc
	)
	if backend == "file" {
		keyDir = filepath.Join(path, "keyring-file")
	} else if backend == "test" {
		keyDir = filepath.Join(path, "keyring-test")
		password = "interstellar"
	}
	if password == "" {
		passFunc = passPrompt(keyName)
	} else {
		passFunc = keyring.FixedStringPrompt(password)
	}
	ring, err := keyring.Open(keyring.Config{
		ServiceName:      "interstellar",
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		FileDir:          keyDir,
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
