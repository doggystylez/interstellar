package keys

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func ensureAddrbook(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			panic(err)
		}
	}
	file := filepath.Join(path, "addrbook.json")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		err := os.WriteFile(file, []byte("{}"), 0700)
		if err != nil {
			panic(err)
		}
	}
}

func SaveToAddrbook(name string, address string, chainId string, path string, internal bool) (err error) {
	var (
		entries *[]Entry
		exists  bool
	)
	book := LoadAddrbook(path)
	if internal {
		entries = &book.InternalAccounts
	} else {
		entries = &book.WithdrawAddresses
	}
	for _, entry := range *entries {
		if entry.Name == name {
			exists = true
			entry.Accounts[chainId] = Account{Address: address}
			break
		}
	}
	if !exists {
		addresses := make(map[string]Account, 1)
		addresses[chainId] = Account{Address: address}
		*entries = append(*entries, Entry{Name: name, Accounts: addresses})
	}
	jsonData, err := json.Marshal(book)
	if err != nil {
		return
	}
	err = os.WriteFile(filepath.Join(path, "addrbook.json"), jsonData, 0700)
	return
}

func SaveAccountInfo(name string, address string, chainId string, accNum uint64, seqNum uint64, path string) (err error) {
	var exists bool
	book := LoadAddrbook(path)
	entries := &book.InternalAccounts
	for _, entry := range *entries {
		if entry.Name == name {
			exists = true
			entry.Accounts[chainId] = Account{
				Address: address,
				AccNum:  accNum,
				SeqNum:  seqNum,
			}
			break
		}
	}
	if !exists {
		addresses := make(map[string]Account, 1)
		addresses[chainId] = Account{
			Address: address,
			AccNum:  accNum,
			SeqNum:  seqNum,
		}
		*entries = append(*entries, Entry{Name: name, Accounts: addresses})
	}
	jsonData, err := json.Marshal(book)
	if err != nil {
		return
	}
	err = os.WriteFile(filepath.Join(path, "addrbook.json"), jsonData, 0700)
	return
}

func LoadAddrbook(path string) (book Addrbook) {
	ensureAddrbook(path)
	f, err := os.Open(filepath.Join(path, "addrbook.json"))
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return
}

func LoadAddress(name string, chainId string, path string, internal bool) string {
	var entries []Entry
	book := LoadAddrbook(path)
	if internal {
		entries = book.InternalAccounts
	} else {
		entries = book.WithdrawAddresses
	}
	for _, entry := range entries {
		if entry.Name == name {
			return entry.Accounts[chainId].Address
		}
	}
	return ""
}

func AddressExists(name string, chainId string, path string, internal bool) bool {
	if LoadAddress(name, chainId, path, internal) == "" {
		return false
	} else {
		return true
	}
}
