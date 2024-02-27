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
	"github.com/progpjs/progpAPI/v2"
)

func jsSamplesGetReturnValues(group *progpAPI.FunctionGroup) {
	group.AddFunction("testString", "JsTestString", JsTestString)
	group.AddFunction("testFloat64", "JsTestFloat64", JsTestFloat64)
	group.AddFunction("testFloat32", "JsTestFloat32", JsTestFloat32)
	group.AddFunction("testInt32", "JsTestInt32", JsTestInt32)
	group.AddFunction("testInt64", "JsTestInt64", JsTestInt64)
	group.AddFunction("testInt", "JsTestInt", JsTestInt)
	group.AddFunction("testBool", "JsTestBool", JsTestBool)

	group.AddFunction("testStructRef", "JsTestStructRef", JsTestStructRef)
	group.AddFunction("testStructPtr", "JsTestStructPtr", JsTestStructPtr)
	group.AddFunction("testStringArray", "JsTestStringArray", JsTestStringArray)
	group.AddFunction("testByteArray", "JsTestByteArray", JsTestByteArray)
}

type TestStructA struct {
	Name     string
	Forename string
}

//region Sample get / return value

func JsTestByteArray(value []byte) []byte {
	return value
}

func JsTestStringArray(value []string) []string {
	return value
}

func JsTestStructRef(value TestStructA) TestStructA {
	return value
}

func JsTestStructPtr(value *TestStructA) *TestStructA {
	return value
}

func JsTestString(value string) string {
	return value
}

func JsTestFloat64(value float64) float64 {
	return value
}

func JsTestFloat32(value float32) float32 {
	return value
}

func JsTestInt32(value int32) int32 {
	return value
}

func JsTestInt64(value int64) int64 {
	return value
}

func JsTestInt(value int) int {
	return value
}

func JsTestBool(value bool) bool {
	return value
}

//endregion
