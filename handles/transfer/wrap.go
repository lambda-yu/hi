package transfer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

func ClientWrapHandle(hdl func(*cli.Context, ethclient.Client, transferData) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		var data transferData
		err := parseArg(c, &data)
		if err != nil {
			return err
		}

		nodeUrl := c.String("url")
		client, err := ethclient.Dial(nodeUrl)
		if err != nil {
			return err
		}

		loop := c.Int("loop")
		if loop == 0 {
			for {
				err := hdl(c, *client, data)
				if err != nil {
					return fmt.Errorf("run transfer error: %s", err)
				}
			}
		} else {
			for i := 0; i < loop; i++ {
				err := hdl(c, *client, data)
				if err != nil {
					return fmt.Errorf("run transfer error: %s, success times: %d/%d", err, i, loop)
				}
			}
		}

		return nil
	}
}
