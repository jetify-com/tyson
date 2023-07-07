package interpreter

import (
	"bytes"
	"os"
	"strings"
	"text/scanner"

	"github.com/evanw/esbuild/pkg/api"
)

var tsonTransform = api.Plugin{
	Name: "tsonTransform",
	Setup: func(build api.PluginBuild) {
		build.OnLoad(
			api.OnLoadOptions{Filter: `\.tson$`},
			loadTSON,
		)
	},
}

func loadTSON(args api.OnLoadArgs) (api.OnLoadResult, error) {
	original, err := os.ReadFile(args.Path)
	if err != nil {
		return api.OnLoadResult{}, err
	}

	offset := findImplicitExport(original)
	var builder strings.Builder

	if offset != -1 {
		builder.Write(original[:offset])
		builder.WriteString("export default ")
		builder.Write(original[offset:])
	} else {
		builder.Write(original)
	}

	result := builder.String()
	return api.OnLoadResult{
		Contents: &result,
		Loader:   api.LoaderTS,
	}, nil
}

// If there are no exports, but there is an top-level object, we identify it
// as an object that should be implicitly exported.
func findImplicitExport(data []byte) int {
	buf := bytes.NewReader(data)
	var tokenizer scanner.Scanner
	tokenizer.Init(buf)
	tokenizer.Error = func(_ *scanner.Scanner, _ string) {} // ignore errors

	var offset = -1
	nestingLevel := 0
	existingObject := false
	for tok := tokenizer.Scan(); tok != scanner.EOF; tok = tokenizer.Scan() {
		switch token := tokenizer.TokenText(); token {
		case "{":
			// We found a top-level object:
			if nestingLevel == 0 {
				if !existingObject {
					// This is the first one we find, so save the offset as we might want to
					// implicitly export it.
					offset = tokenizer.Offset
					existingObject = true
				} else {
					// If we've found more than one top-level object, we don't want to implicitly
					// export any of them.
					return -1
				}
			}
			nestingLevel++
		case "}":
			nestingLevel--
		default:
			// We've run into another expression, so we don't want to implicitly export anything.
			if nestingLevel == 0 {
				return -1
			}
		}
	}
	return offset
}
