package client

import (
	"errors"
	"github.com/spf13/viper"
	"go.starlark.net/starlark"
)

var (
	thread                 *starlark.Thread
	filterFunction         starlark.Value
	filterFunctionProvided bool
)

//setupFilter creates the Starlark filter function if provided
func setupFilter() error {
	ff := viper.GetString("connect.filter-function")
	if ff == "" {
		return nil
	}

	// Execute Starlark program in a file.
	thread = &starlark.Thread{Name: "filter thread"}
	globals, err := starlark.ExecFile(thread, ff, nil, nil)
	if err != nil {
		return err
	}

	if _, ok := globals["filter"]; !ok {
		return errors.New("function filter not found")
	}

	// Retrieve a module global.
	filterFunction = globals["filter"]

	filterFunctionProvided = true

	return nil
}
