package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	version = "0.0.1"
)

type option struct {
	configPath string
}

var opt = &option{}

func main() {
	app := cli.NewApp()
	app.Name = "mock"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Value:       "config.json",
			Usage:       "config path",
			Destination: &opt.configPath,
		},
	}

	app.Action = func(c *cli.Context) error {
		if err := run(opt); err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	}

	app.Run(os.Args)
}
