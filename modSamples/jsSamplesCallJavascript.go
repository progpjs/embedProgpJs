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

// === About auto disposing of C++ resources ===
//
// To know: when the GC dispose a progpAPI.ScriptFunction then it automatically
// dispose the C++ resource linked to this object. But it can take a lot of time beforing
// the GC execute and it's incompatible will high load scenario where.
//
// It why by default the C++ resource are disposed once the first call done.
// If you want to use the same ref more than one time, you must call "myJsFunction.KeepAlive()".
// (it must be done only one time).

func jsSamplesCallJavascript(group *progpAPI.FunctionGroup) {
	group.AddFunction("testCallbackWithString", "JsCallbackWithString", JsCallbackWithString)
	group.AddFunction("testCallbackWithBool", "JsCallbackWithBool", JsCallbackWithBool)
	group.AddFunction("testCallbackWithDouble", "JsCallbackWithDouble", JsCallbackWithDouble)
	group.AddFunction("testCallbackWithError", "JsCallbackWithError", JsCallbackWithError)
	group.AddFunction("testCallbackWithoutValues", "JsCallbackWithoutValues", JsCallbackWithoutValues)
	group.AddFunction("testCallbackWithArrayBuffer", "JsCallbackWithArrayBuffer", JsCallbackWithArrayBuffer)
	group.AddFunction("testCallbackWithStringBuffer", "JsCallbackWithStringBuffer", JsCallbackWithStringBuffer)
}

func JsCallbackWithArrayBuffer(fct progpAPI.JsFunction, value []byte) {
	fct.CallWithArrayBuffer2(value)
}

func JsCallbackWithStringBuffer(fct progpAPI.JsFunction, value string) {
	// Here it's an optimization avoiding converting from buffer to string before calling javascript.
	asBytes := []byte(value)
	fct.CallWithStringBuffer2(asBytes)
}

func JsCallbackWithString(fct progpAPI.JsFunction, value string) {
	fct.CallWithString2(value)
}

func JsCallbackWithBool(fct progpAPI.JsFunction, value bool) {
	fct.CallWithBool2(value)
}

func JsCallbackWithDouble(fct progpAPI.JsFunction, value float64) {
	fct.CallWithDouble2(value)
}

func JsCallbackWithError(fct progpAPI.JsFunction, errMessage string) {
	fct.CallWithError(errors.New(errMessage))
}

func JsCallbackWithoutValues(fct progpAPI.JsFunction) {
	fct.CallWithUndefined()
}
