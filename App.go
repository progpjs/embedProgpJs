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

package main

import (
	"github.com/progpjs/embedProgpJs/v2/modSamples"
	"github.com/progpjs/modules/v2/modCore"
	"github.com/progpjs/modules/v2/modHttp"
	"github.com/progpjs/modules/v2/modNodeJs"
	"github.com/progpjs/modules/v2/modReact"
)

const MustDebug = false

func main() {
	awaiter := bootstrapProgpJS("index.ts", MustDebug, nil, RegisterMyModules)

	// Will wait until the VM can exit.
	//
	// The VM can't exit if there is background task remaining (ex: a webserver).
	// If you don't call awaiter then your app will quit immediately and your server
	// will not be able to execute.
	//
	awaiter()
}

// RegisterMyModules registers our all ProgpJS modules.
func RegisterMyModules() {
	// Required core modules.
	//
	modCore.InstallProgpJsModule()
	modNodeJs.InstallProgpJsModule()

	// Optional core modules.
	//
	modReact.InstallProgpJsModule()
	modHttp.InstallProgpJsModule()
	//progpJsonDB.InstallProgpJsModule()

	// Our own modules.
	modSamples.InstallModule()
}
