package config

import (
	"github.com/urfave/cli/v2"
)

// create account util flag
var (
	AccountSavePathFlag = &cli.StringFlag{
		Name:    "save",
		Aliases: []string{"s"},
		Usage:   "Save path for account file",
	}
	AccountNumberFlag = &cli.Int64Flag{
		Name:     "number",
		Aliases:  []string{"n"},
		Usage:    "Number of create accounts",
		Required: true,
	}
)

var AccountFlags = []cli.Flag{
	AccountNumberFlag,
	AccountSavePathFlag,
}

// eth client flag
var (
	NodeURL = &cli.StringFlag{
		Name:     "url",
		Usage:    "Block chain node host. ex: http://ip:port",
		Required: true,
	}
	ChainID = &cli.Int64Flag{
		Name:     "chainid",
		Usage:    "Block chain node id",
		Required: true,
	}
)

// transfer util flag
var (
	TransferRecipientFlag = &cli.StringSliceFlag{
		Name:  "recipient",
		Usage: "receiver address",
	}
	TransferRecipientFileFlag = &cli.PathFlag{
		Name:  "recipients",
		Usage: "receiver address",
	}
	TransferSenderFlag = &cli.StringFlag{
		Name:     "sender",
		Usage:    "sender private",
		Required: true,
	}
	TransferSendAccountFileFlag = &cli.PathFlag{
		Name:     "sendAccounts",
		Usage:    "path of send account file",
		Required: true,
	}
	TransferRecipientAccountFileFlag = &cli.PathFlag{
		Name:     "recipientAccounts",
		Usage:    "path of recipient account file",
		Required: true,
	}
	TransferAmountFlag = &cli.StringFlag{
		Name:     "amount",
		Usage:    "amount of the transfer",
		Required: true,
	}
	TransferLoopTimesFlag = &cli.IntFlag{
		Name:  "loop",
		Usage: "transfer execute number of cycle, if 0 then infinite loop",
		Value: 1,
	}
	TransferGasLimitFlag = &cli.Uint64Flag{
		Name:  "gaslimit",
		Value: 21000,
	}
)

var TransferFlags = []cli.Flag{
	NodeURL,
	ChainID,
	TransferRecipientFlag,
	TransferSenderFlag,
	TransferAmountFlag,
	TransferGasLimitFlag,
	TransferLoopTimesFlag,
	TransferRecipientFileFlag,
}

var TransferMany2ManyFlags = []cli.Flag{
	NodeURL,
	ChainID,
	TransferSendAccountFileFlag,
	TransferRecipientAccountFileFlag,
	TransferAmountFlag,
	TransferGasLimitFlag,
	TransferLoopTimesFlag,
}
