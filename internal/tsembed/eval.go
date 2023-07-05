package tsembed

import "github.com/dop251/goja"

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
