package file

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"strings"
)

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func (a Account) PrivateHex() string {
	return hexutil.Encode(crypto.FromECDSA(a.PrivateKey))[2:]
}

func Save(path string, keys []Account) error {
	var buffer bytes.Buffer
	for _, key := range keys {
		msg := fmt.Sprintf("%s %s\n", key.Address, key.PrivateHex())
		buffer.Write([]byte(msg))
	}

	return ioutil.WriteFile(path, buffer.Bytes()[:buffer.Len()-1], 0644)
}

func Read(path string) ([]Account, error) {
	var keys []Account
	sendAccountBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read send account error: %s", err)
	}
	sendAccounts := strings.Split(string(sendAccountBytes), "\n")
	for i := 0; i < len(sendAccounts); i++ {
		sAccount := strings.Split(sendAccounts[i], " ")
		sP, err := crypto.HexToECDSA(sAccount[1])
		if err != nil {
			return nil, fmt.Errorf("read send account private error: %s", err)
		}
		keys = append(keys, Account{common.HexToAddress(sAccount[0]), sP})
	}
	return keys, nil
}
