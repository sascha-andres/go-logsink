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
