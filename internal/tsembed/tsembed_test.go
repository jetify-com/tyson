package tsembed

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "export object",
			input: `
				export default {
					string_field: "string",
					number_field: 123,
					boolean_field: true,
					array_field: [1, 2, 3],
					object_field: {
						name: "object"
					}
				}
			`,
			expected: `
				{
					"string_field": "string",
					"number_field": 123,
					"boolean_field": true,
					"array_field": [1, 2, 3],
					"object_field": {
						"name": "object"
					}
				}
			`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "input.ts")
			err := os.WriteFile(path, []byte(tt.input), 0o644)
			assert.NoError(t, err)
			val, err := Eval(path, Options{})
			assert.NoError(t, err)
			jsonBytes, err := json.Marshal(val)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(jsonBytes))
		})
	}
}
