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
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "times",
			Usage:       "Print messages multiple times.",
			Description: "Print messages multiple times.",
			ArgsUsage:   "[message]",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n,times",
					Value: 1,
				},
			},
			Action: func(ctx *cli.Context) error {
				msg := ctx.Args().First()
				if len(msg) == 0 {
					msg = "hello"
				}
				times := ctx.Int("times")
				for i := 0; i < times; i++ {
					fmt.Println(msg)
				}
				return nil
			},
		},
		cli.Command{
			Name:        "reverse",
			Usage:       "Print reverse ordered message.",
			Description: "Print reverse ordered message.",
			Action: func(ctx *cli.Context) error {
				msg := []rune(ctx.Args().First())
				for i := 0; i < len(msg); i++ {
					fmt.Print(string(msg[len(msg)-i-1]))
				}
				fmt.Println()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
