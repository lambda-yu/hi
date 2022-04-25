package account

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"path/filepath"
	"utils/config"
	"utils/pkg/file"
)

// HandleGenerateCmd handle for generate account
func HandleGenerateCmd(c *cli.Context) error {
	var keys []file.Account
	number := c.Int64("number")
	if number == 0 {
		number = 1
	}

	keys, err := generate(number)
	if err != nil {
		return err
	}

	pathStr := c.String("save")
	// not set path , then use default
	if pathStr == "" {
		pathStr = filepath.Join(config.DefaultSavePath, config.DefaultSaveName)
	}

	s, err := os.Stat(pathStr)
	if err == nil && s.IsDir() {
		pathStr = path.Join(pathStr, config.DefaultSaveName)
	}

	return file.Save(pathStr, keys)
}

func generate(num int64) ([]file.Account, error) {
	var accounts []file.Account
	for i := int64(0); i < num; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("error casting public key to ECDSA")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		accounts = append(accounts, file.Account{Address: common.HexToAddress(address), PrivateKey: privateKey})
	}
	return accounts, nil
}
