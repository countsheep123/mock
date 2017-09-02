package main

import (
	"github.com/countsheep123/mock"
)

func run(opt *option) error {
	handlers, err := load(opt.configPath)
	if err != nil {
		return err
	}

	mock.Serve(handlers)
	return nil
}
