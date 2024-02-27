/*
 * (C) Copyright 2024 Johan Michel PIQUET, France (https://johanpiquet.fr/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package modSamples

import (
	"errors"
	"github.com/progpjs/progpAPI/v2"
)

func jsSamplesThrowError(group *progpAPI.FunctionGroup) {
	group.AddFunction("testThrowError", "JsThrowError", JsThrowError)
	group.AddFunction("testThrowErrorP1", "JsThrowErrorP1", JsThrowErrorP1)
	group.AddFunction("testThrowErrorP2", "JsThrowErrorP2", JsThrowErrorP2)
}

func JsThrowError(value string) error {
	if value == "boom" {
		return errors.New("big boom")
	}

	return nil
}

func JsThrowErrorP1(value string) (error, string) {
	if value == "boom" {
		return errors.New("big boom"), ""
	}

	return nil, value
}

func JsThrowErrorP2(value string) (string, error) {
	if value == "boom" {
		return "", errors.New("big boom")
	}

	return value, nil
}
