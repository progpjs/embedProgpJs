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

import "github.com/progpjs/progpAPI/v2"

func jsSamplesSharedResources(group *progpAPI.FunctionGroup) {
	group.AddFunction("testReceiveSharedResource", "JsTestReceiveSharedResource", JsTestReceiveSharedResource)
	group.AddFunction("testReturnSharedResource", "JsTestReturnSharedResource", JsTestReturnSharedResource)
}

func JsTestReceiveSharedResource(sr *progpAPI.SharedResource) {
	println("Received shared resource: ", sr.GetId(), " - Value: ", sr.Value.(string))
}

func JsTestReturnSharedResource(rc *progpAPI.SharedResourceContainer) *progpAPI.SharedResource {
	sr := rc.NewSharedResource("my resource", nil)
	return sr
}
