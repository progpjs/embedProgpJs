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
	"github.com/progpjs/progpjs/v2"
)

// === About functions callers ===
//
// Functions callers allows to call a javascript function from Go.
// It generates Go and C++ code which allows calling javascript functions the way you want.
// progpAPI.JsFunction offers some possibilities, but are limited and you will probably require
// to use a "function caller" in order to call your custom javascript functions.

func jsSamplesFunctionCaller(group *progpAPI.FunctionGroup) {
	// region Sample : calling (string, string)

	// Here we create a custom caller.
	//
	// It uses an objet as his only parameter, which act as a sample
	// allowing the engine to know how to call our function. This sample
	// object require a method named "Call" which first parameter is of type progpAPI.JsFunction
	// and where the others parameters are a mirror of our javascript function parameters.
	//
	// Here for our sample we use (string, string) as parameters, but our can
	// us any combination with string, int, float64, bool, []byte.
	//
	// We also use an interface, which will allow to replace our fake objet by
	// an object implementing our interface. The only reason of this interface is that:
	// being able to replace our fake object by the real one. Without an interface the
	// Go compile will not accept the change.
	//
	//
	gMyJavascriptFunctionCaller = progpjs.GetFunctionCaller(&implStringStringCaller{}).(StringStringCaller)

	// Here it's a function allowing to test our caller.
	// From javascript we can do : testStringStringCaller((name, forename) => console.log("Hello ", name, " ", forename));
	//
	group.AddFunction("testStringStringCaller", "JsTestStringStringCaller", JsTestStringStringCaller)

	// About return params ... the engine doesn't support return params.
	// Why? Mainly because calling javascript function is asynchronous,
	// V8 add the call in a task list an process the call once he finished
	// his previous call. It's why javascript call can only be asynchronous.
	// There is possibly hack to avoid it, but we think it's not useful in real life apps.

	// endregion

	//region All supported types

	// Here is a demo with all supported types.
	// To test it:
	//
	/*
		testAllSupportedTypesCaller((pString, pBool, pNumber, pArrayBuffer, pStringBuffer) => {
		  console.log("pString=", pString);
		  console.log("pBool=", pBool);
		  console.log("pNumber=", pNumber);
		  console.log("pArrayBuffer=", progpBufferToString(pArrayBuffer));
		  console.log("pStringBuffer=", pStringBuffer);
		})
	*/

	getAllSupportedTypesCaller = progpjs.GetFunctionCaller(progpjs.GetFunctionCaller(&implAllSupportedTypesCaller{})).(AllSupportedTypesCaller)
	group.AddFunction("testAllSupportedTypesCaller", "JsTestAllSupportedTypesCaller", JsTestAllSupportedTypesCaller)

	//endregion
}

func JsTestStringStringCaller(callback progpAPI.JsFunction) {
	// Call the javascript function with our custom caller.
	gMyJavascriptFunctionCaller.Call(callback, "John", "Doe")
}

func JsTestAllSupportedTypesCaller(callback progpAPI.JsFunction) {
	getAllSupportedTypesCaller.Call(
		callback,
		"my string",
		true,
		123.456,
		[]byte("array buffer sample"),
		progpAPI.StringBuffer("here it's a []byte but for js it's a string"),
	)
}

//region Javascript call of type (string, string)

var gMyJavascriptFunctionCaller StringStringCaller

// StringStringCaller is the interface used to build a javascript caller
// able to call function requiring two string as parameters.
type StringStringCaller interface {
	Call(js progpAPI.JsFunction, param1 string, param2 string)
}

// implStringStringCaller si the default implementation for StringStringCaller.
// This default version must call the dynamic caller, without that you won't be able
// to use the dynamic mode, which avoid to generate compiled code.
type implStringStringCaller struct {
}

func (*implStringStringCaller) Call(js progpAPI.JsFunction, param1 string, param2 string) {
	js.DynamicFunctionCaller(param1, param2)
}

//endregion

//region Javascript call with all supported types

var getAllSupportedTypesCaller AllSupportedTypesCaller

type AllSupportedTypesCaller interface {
	Call(js progpAPI.JsFunction, testString string, testBool bool, testFloat64 float64, testArrayBuffer []byte, testStringBuffer progpAPI.StringBuffer)
}

type implAllSupportedTypesCaller struct {
}

func (*implAllSupportedTypesCaller) Call(js progpAPI.JsFunction, testString string, testBool bool, testFloat64 float64, testArrayBuffer []byte, testStringBuffer progpAPI.StringBuffer) {
	js.DynamicFunctionCaller(testString, testBool, testFloat64, testArrayBuffer, testStringBuffer)
}

//endregion
