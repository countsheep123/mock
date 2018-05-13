package main

import (
	"encoding/json"
	"os"

	"github.com/countsheep123/mock"
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

func run(opt *option) error {
	handlers, err := load(opt.configPath)
	if err != nil {
		return err
	}

	mock.Serve(handlers)
	return nil
}

func load(filepath string) ([]*mock.Handler, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	var handlers []*mock.Handler
	decoder := json.NewDecoder(f)
	decoder.UseNumber()
	err = decoder.Decode(&handlers)
	if err != nil {
		return nil, err
	}

	return handlers, nil
}
