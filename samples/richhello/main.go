package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "n,times",
			Value: 1,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		msg := ctx.Args().First()
		// msgのデフォルト値設定
		times := ctx.Int("times")
		// times回、msgを標準出力に出力
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
