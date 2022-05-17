package transfer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
	"math/big"
	"sync"
	"time"
	"utils/pkg/file"
)

type transferData struct {
	FromAccounts []file.Account
	ToAccounts   []file.Account
	Value        *big.Int
	GasLimit     uint64
	ChainID      *big.Int
}

func HandleOne2One(c *cli.Context, client ethclient.Client, data transferData) error {
	return transfer2Many(c.Context, client, data)
}

func HandleMany2Many(c *cli.Context, client ethclient.Client, data transferData) error {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for i := 0; i < len(data.FromAccounts); i++ {
		go transfer(c.Context, client, wg, data.ToAccounts[i].Address, data.FromAccounts[i].PrivateKey, data.Value, data.ChainID, data.GasLimit)
	}
	time.Sleep(time.Second * 1)
	wg.Wait()
	return nil
}

func HandleMany2One(c *cli.Context, client ethclient.Client, data transferData) error {
	return transferMany2One(c.Context, client, data)
}

func parseArg(c *cli.Context, data *transferData) error {
	var ok bool
	var err error

	tos := c.StringSlice("recipient")
	toPath := c.Path("recipientAccounts")

	if len(tos) > 0 {
		for _, to := range tos {
			data.ToAccounts = append(data.ToAccounts, file.Account{Address: common.HexToAddress(to), PrivateKey: nil})
		}
	} else if toPath != "" {
		data.ToAccounts, err = file.Read(toPath)
		if err != nil {
			return fmt.Errorf("read recipient accout error: %s", err)
		}
	}

	senders := c.StringSlice("sender")
	sendPath := c.Path("sendAccounts")

	if len(senders) > 0 {
		for _, sender := range senders {
			fromPivateKey, err := crypto.HexToECDSA(sender)
			if err != nil {
				return fmt.Errorf("get sender error")
			}

			fromPub := fromPivateKey.Public()
			fromPubEcdsa, ok := fromPub.(*ecdsa.PublicKey)
			if !ok {
				return fmt.Errorf("error casting public key to ECDSA")
			}

			FromAddress := crypto.PubkeyToAddress(*fromPubEcdsa)
			data.FromAccounts = append(data.FromAccounts, file.Account{FromAddress, fromPivateKey})
		}
	} else {
		data.FromAccounts, err = file.Read(sendPath)
		if err != nil {
			return fmt.Errorf("read recipient accout error: %s", err)
		}
	}

	data.Value, ok = (&big.Int{}).SetString(c.String("amount"), 10)
	if !ok {
		return fmt.Errorf("get amount error")
	}

	data.GasLimit = c.Uint64("gaslimit")

	data.ChainID = big.NewInt(c.Int64("chainid"))

	// if many to many then need account file
	send := c.Path("sendAccounts")
	if send != "" {
		data.FromAccounts, err = file.Read(send)
		if err != nil {
			return fmt.Errorf("read recipient accout error: %s", err)
		}
	}

	recipient := c.Path("recipientAccounts")
	if recipient != "" {
		data.ToAccounts, err = file.Read(recipient)
		if err != nil {
			return fmt.Errorf("read recipient accout error: %s", err)
		}
	}

	return nil
}

func transfer2Many(ctx context.Context, client ethclient.Client, data transferData) error {
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	singer := types.NewLondonSigner(data.ChainID)

	var successAddress []common.Address
	for _, to := range data.ToAccounts {
		nonce, err := client.PendingNonceAt(ctx, data.FromAccounts[0].Address)
		if err != nil {
			return err
		}
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			Value:    data.Value,
			To:       &to.Address,
			GasPrice: gasPrice,
			Gas:      data.GasLimit,
			Data:     nil,
		})

		signedTx, err := types.SignTx(tx, singer, data.FromAccounts[0].PrivateKey)

		err = client.SendTransaction(ctx, signedTx)
		if err != nil {
			return fmt.Errorf("[%s] send transaction error: %s, send success %d/%d, success address: %s", to, err, len(successAddress), len(data.ToAccounts), successAddress)
		}
		successAddress = append(successAddress, to.Address)
	}
	return nil
}

func transferMany2One(ctx context.Context, client ethclient.Client, data transferData) error {
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	singer := types.NewLondonSigner(data.ChainID)

	var successAddress []common.Address
	for _, From := range data.FromAccounts {
		nonce, err := client.PendingNonceAt(ctx, From.Address)
		if err != nil {
			return err
		}
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			Value:    data.Value,
			To:       &data.ToAccounts[0].Address,
			GasPrice: gasPrice,
			Gas:      data.GasLimit,
			Data:     nil,
		})

		signedTx, err := types.SignTx(tx, singer, From.PrivateKey)

		err = client.SendTransaction(ctx, signedTx)
		if err != nil {
			return fmt.Errorf("[%s] send transaction error: %s, send success %d/%d, success address: %s", From, err, len(successAddress), len(data.FromAccounts), successAddress)
		}
		successAddress = append(successAddress, From.Address)
	}
	return nil
}

func transfer(ctx context.Context, client ethclient.Client, wg *sync.WaitGroup, to common.Address, sender *ecdsa.PrivateKey, value *big.Int, chainID *big.Int, gaslimit uint64) {
	wg.Add(1)
	defer wg.Done()
	fromPub := sender.Public()
	fromPubEcdsa, ok := fromPub.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("sender error")
		return
		//fmt.Errorf("[%s] send transaction error: %s, send success %d/%d, success address: %s")
	}

	fromAddress := crypto.PubkeyToAddress(*fromPubEcdsa)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		fmt.Println(err)
		return

	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println(err)
		return

	}

	singer := types.NewLondonSigner(chainID)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Value:    value,
		To:       &to,
		GasPrice: gasPrice,
		Gas:      gaslimit,
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, singer, sender)
	if err != nil {
		fmt.Println(err)
		return

	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		fmt.Println(err)
		return

	}
}
