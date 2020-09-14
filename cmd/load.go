package cmd

import (
	"github.com/go-openapi/loads"
)

func load(filename string) (*loads.Document, error) {
	d, err := loads.JSONSpec(filename)
	if err != nil {
		return nil, err
	}
	return d, nil
}
