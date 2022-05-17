package account

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
)

func HandleGenKeystore(c *cli.Context) error {
	return nil
}

func createKs(password string, savePath string) error {
	// new keystore
	ks := keystore.NewKeyStore(savePath, keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount(password)
	if err != nil {
		return err
	}
	return nil
}

func importKs(file string, password string) (*keystore.KeyStore, error) {
	// new a temp keystore for import
	ks := keystore.NewKeyStore("./temp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// import
	_, err = ks.Import(jsonBytes, password, password)
	if err != nil {
		return nil, err
	}

	// rm temp
	if err := os.Remove(ks.Accounts()[0].URL.Path); err != nil {
		return nil, err
	}
	return ks, nil
}
