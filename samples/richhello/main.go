package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "richhello"
	app.Version = "1.0.0"
	app.Usage = "Print messages multiple times."
	app.Description = "Print messages multiple times."

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "n,times",
			Value: 1,
		},
	}
	app.ArgsUsage = "[message]"
	app.Action = func(ctx *cli.Context) error {
		msg := ctx.Args().First()
		if len(msg) == 0 {
			msg = "hello"
		}
		times := ctx.Int("times")
		for i := 0; i < times; i++ {
			fmt.Println(msg)
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
