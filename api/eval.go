package api

import (
	"encoding/json"

	"go.jetpack.io/tyson/internal/tsembed"
)

func Eval(inputPath string) ([]byte, error) {
	v, err := tsembed.Eval(inputPath)

	if err != nil {
		return nil, err
	}

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
