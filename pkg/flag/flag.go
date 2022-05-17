package flag

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"strings"
	"syscall"
)

type privateValue struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func NewPrivateKeyValue() *privateValue {
	return &privateValue{}
}

func (p *privateValue) Set(value string) error {
	var err error
	p.PrivateKey, err = crypto.HexToECDSA(value)
	if err != nil {
		return err
	}
	fromPub := p.PrivateKey.Public()
	fromPubEcdsa, ok := fromPub.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	p.Address = crypto.PubkeyToAddress(*fromPubEcdsa)
	return nil
}

func (p *privateValue) String() string {
	private := crypto.FromECDSA(p.PrivateKey)
	return hexutil.Encode(private)
}

func (p *privateValue) Get() any {
	return *p
}

func flagFromEnvOrFile(envVars []string, filePath string) (val string, ok bool) {
	for _, envVar := range envVars {
		envVar = strings.TrimSpace(envVar)
		if val, ok := syscall.Getenv(envVar); ok {
			return val, true
		}
	}
	for _, fileVar := range strings.Split(filePath, ",") {
		if fileVar != "" {
			if data, err := ioutil.ReadFile(fileVar); err == nil {
				return string(data), true
			}
		}
	}
	return "", false
}
