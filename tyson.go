package tyson

import (
	"go.jetpack.io/tyson/api"
)

// Eval evaluates a tson file and returns the result as a JSON-encoded byte slice.
func Eval(tsonPath string) ([]byte, error) {
	return api.Eval(tsonPath)
}

// Unmarshal is a convenience function that first evaluates the given TSON file,
// and then unmarshals the result into the given go struct.
// Internally it unmarshals using json.Unmarshal, so the behavior is the same.
func Unmarshal(tsonPath string, v any) error {
	return api.Unmarshal(tsonPath, v)
}
