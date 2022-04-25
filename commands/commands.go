package commands

import (
	"github.com/urfave/cli/v2"
	"utils/config"
	"utils/handles/account"
	"utils/handles/transfer"
)

var (
	AccountGenerateCommand = &cli.Command{
		Name:   "GenAccount",
		Flags:  config.AccountFlags,
		Action: account.HandleGenerateCmd,
	}

	TransferCommand = &cli.Command{
		Name: "Transfer",
		Subcommands: []*cli.Command{
			one2OneCommand,
			one2ManyCommand,
			many2ManyCommand,
			many2OneCommand,
		},
	}
)

// transfer subcommand
var (
	one2ManyCommand = &cli.Command{
		Name:   "OneToOne",
		Flags:  config.TransferOne2ManyFlags,
		Action: transfer.ClientWrapHandle(transfer.HandleOne2One),
	}
	one2OneCommand = &cli.Command{
		Name:   "OneToMany",
		Flags:  config.TransferOne2OneFlags,
		Action: transfer.ClientWrapHandle(transfer.HandleOne2One),
	}
	many2ManyCommand = &cli.Command{
		Name:   "ManyToMany",
		Flags:  config.TransferMany2ManyFlags,
		Action: transfer.ClientWrapHandle(transfer.HandleMany2Many),
	}
	many2OneCommand = &cli.Command{
		Name:   "ManyToOne",
		Flags:  config.TransferMany2OneFlags,
		Action: transfer.ClientWrapHandle(transfer.HandleMany2One),
	}
)
