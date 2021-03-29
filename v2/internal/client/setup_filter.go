// Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
