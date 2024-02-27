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

// === About async functions ===
//
// Async functions must end with "Async" (only for the Go function)
// and be called with AddAsyncFunction. This will add a special flag,
// without which the VM can exit before the function end.
//
// You must be warning of that: doing async without an async function
// works, be can have behaviors since the system don't know it's async
// and that he must wait for a response.

func jsSamplesCallAsync(group *progpAPI.FunctionGroup) {
	group.AddAsyncFunction("testAsync", "JsTestAsync", JsTestAsync)
}

func JsTestAsync(name string, callback progpAPI.JsFunction) {
	progpAPI.SafeGoRoutine(func() {
		// Add a pause for sample
		progpAPI.PauseMs(100)
		callback.CallWithString2("hello " + name)
	})
}
