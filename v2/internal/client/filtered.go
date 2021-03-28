package client

import "go.starlark.net/starlark"

//filtered returns true if a filter function is provided and it evaluates to true
func filtered(line string) bool {
	if !filterFunctionProvided {
		return false
	}
	v, err := starlark.Call(thread, filterFunction, starlark.Tuple{starlark.String(line)}, nil)
	if err != nil {
		return false
	}
	return v == starlark.True
}
