package api

import "encoding/json"

func Unmarshal(tsonPath string, v any) error {
	bytes, err := Eval(tsonPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}
