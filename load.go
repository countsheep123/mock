package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func load(filepath string) ([]*handler, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var handlers []*handler
	decoder := json.NewDecoder(bytes.NewReader(file))
	decoder.UseNumber()
	err = decoder.Decode(&handlers)
	if err != nil {
		return nil, err
	}

	return handlers, nil
}
