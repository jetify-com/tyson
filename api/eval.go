package api

import (
	"encoding/json"

	"go.jetpack.io/tyson/internal/interpreter"
)

func Eval(inputPath string) ([]byte, error) {
	v, err := interpreter.Eval(inputPath)
	if err != nil {
		return nil, err
	}

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
