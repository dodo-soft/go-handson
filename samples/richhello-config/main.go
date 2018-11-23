package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Message string `yaml:"message"`
	Times   int    `yaml:"times"`
}

func main() {
	app := cli.NewApp()
	app.Name = "richhello"
	app.Version = "1.0.0"
	app.Usage = "Print messages multiple times."
	app.Description = "Print messages multiple times."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "c,config",
			EnvVar: "HELLO_CONFIG",
		},
	}
	app.ArgsUsage = "[message]"
	app.Action = func(ctx *cli.Context) error {
		configFile := ctx.String("config")

		// ファイルをオープンして
		f, err := os.Open(configFile)
		if err != nil {
			return err
		}
		defer f.Close()
		// 全データを[]byte型で読み込んで
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		// Config構造体にマッピングする
		var config *Config
		if err := yaml.Unmarshal(b, &config); err != nil {
			return err
		}

		msg := config.Message
		if len(msg) == 0 {
			msg = "hello"
		}
		times := config.Times
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
