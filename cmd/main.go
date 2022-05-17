package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"utils/commands"
	"utils/config"
	"utils/pkg/flag"
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
	//app()
	app := &cli.App{
		Name:    config.Name,
		Usage:   config.Usage,
		Version: config.Version,
		Flags: []cli.Flag{
			&flag.PrivateKeyFlag{
				Name: "private",
			},
			&cli.StringFlag{
				Name: "TEST",
			},
		},
		Action: func(c *cli.Context) error {
			v := c.Value("TEST")
			fmt.Println(v)
			f := c.Value("private")
			fmt.Println(f)
			return nil
		},
	}
	fmt.Println(os.Args)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
