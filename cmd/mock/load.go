package main

import (
	"encoding/json"
	"os"

	"github.com/countsheep123/mock"
)

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
