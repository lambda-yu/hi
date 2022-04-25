package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"utils/commands"
	"utils/config"
)

func app() {
	app := &cli.App{
		Name:    config.Name,
		Usage:   config.Usage,
		Version: config.Version,
		Commands: []*cli.Command{
			commands.AccountGenerateCommand,
			commands.TransferCommand,
		},
	}

	app.Commands = append(app.Commands)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app()
}
