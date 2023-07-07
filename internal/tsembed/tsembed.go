package tsembed

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"go.jetpack.io/tyson/msgerror"
)

type Options struct {
	Plugins []api.Plugin
}

func Eval(entrypoint string, opts Options) (goja.Value, error) {
	bundle, err := Build(entrypoint, opts)
	if err != nil {
		return nil, err
	}
	return evalJS(string(bundle))
}

func evalJS(code string) (goja.Value, error) {
	vm := goja.New()
	_, err := vm.RunString(code)
	if err != nil {
		return nil, err
	}
	globals := vm.Get(globalsName)
	// Return null if the globals variable is not defined.
	if globals == nil || goja.IsNull(globals) || goja.IsUndefined(globals) {
		return nil, nil
	}
	val := globals.ToObject(vm).Get("default")
	// Right now we return a goja value, but this might have to change if we
	// decide to move to V8
	return val, nil
}

// Default tsConfig
const tsConfig = `
{
  "compilerOptions": {
    "allowJs": true,
    "esModuleInterop": true,
    "experimentalDecorators": true,
    "inlineSourceMap": true,
    "isolatedModules": true,
    "module": "esnext",
    "moduleDetection": "force",
    "strict": true,
    "target": "es6",
    "useDefineForClassFields": true
  }
}
`

const globalsName = "globals"

func Build(entrypoint string, opts Options) ([]byte, error) {
	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{entrypoint},

		Bundle:      true,
		Charset:     api.CharsetUTF8,
		GlobalName:  globalsName,
		Plugins:     opts.Plugins,
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
