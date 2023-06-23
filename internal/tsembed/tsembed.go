package tsembed

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"go.jetpack.io/tyson/msgerror"
)

func Eval(entrypoint string) (goja.Value, error) {
	bundle, err := Build(entrypoint)
	if err != nil {
		return nil, err
	}
	return evalJS(string(bundle))
}

// We use the same tsconfig.json as Deno, see:
// https://deno.com/manual@v1.34.3/advanced/typescript/configuration#what-an-implied-tsconfigjson-looks-like
// With the exception of setting target to "es5" instead of "esnext" so that
// we can run the output in engines like GoJa, that only support ES5.
const tsConfig = `
{
  "compilerOptions": {
    "allowJs": true,
    "esModuleInterop": true,
    "experimentalDecorators": true,
    "inlineSourceMap": true,
    "isolatedModules": true,
    "jsx": "react",
    "module": "esnext",
    "moduleDetection": "force",
    "strict": true,
    "target": "es6",
    "useDefineForClassFields": true
  }
}
`

const globalsName = "globals"

func Build(entrypoint string) ([]byte, error) {
	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{entrypoint},

		Bundle:     true,
		Charset:    api.CharsetUTF8,
		GlobalName: globalsName,
		Loader: map[string]api.Loader{
			".tson": api.LoaderTS,
		},
		Platform:    api.PlatformBrowser,
		Target:      api.ES2015, // ES6 == ES2015
		TsconfigRaw: tsConfig,
		Write:       false,
	})

	if len(bundle.Errors) > 0 {
		msg := fmt.Sprintf("%d syntax errors when compiling %s", len(bundle.Errors), entrypoint)
		return nil, msgerror.ErrFromMessages(msg, bundle.Errors)
	}

	if len(bundle.OutputFiles) != 1 {
		return nil, fmt.Errorf("expected 1 output file, got %d", len(bundle.OutputFiles))
	}

	return bundle.OutputFiles[0].Contents, nil
}
