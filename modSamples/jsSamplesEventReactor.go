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

// Here is a more complexe sample showing you how to implement an event / reactor logic
// which means having a javascript function called each time an event occurs in the Go side.
// In this sample the event is triggered every 2 secondes, starting once a reactor (javascript function) is registered.
//
// Some special mechanisms are used here, like automatic resource managing which allows to automatically dispose
// the resources created by our reactor function. It's important since without that our resource are disposed when
// the script exit, which is never the case here du to the fact that we must enable long-running VM functionality.
// Without enabling this functionality our javascript VM exit once the script is finished (and all the executing async).

// Sample usage from javascript:
/*
testRegisterEventReactor(() => {
	// Will print "disposing my first resource" when auto-disposing".
	testCreateResource("my first resource");

	// In real project a resource is something like a file descriptor
	// or a link to the curent http call. Something coming from Go.
	//
	testCreateResource("my second resource");
})
*/

func jsSamplesEventReactor(group *progpAPI.FunctionGroup) {
	group.AddFunction("testCreateResource", "JsTestCreateResource", JsTestCreateResource)
	group.AddFunction("testRegisterEventReactor", "JsTestRegisterEventReactor", JsTestRegisterEventReactor)
}

// Contains the function to execute when our event occurs.
var gEventReactor progpAPI.JsFunction

func triggerMyEvent() {
	if gEventReactor != nil {
		gEventReactor.CallWithUndefined()
	}
}

// JsTestRegisterEventReactor register the function reacting to our event.
// Once done its start emitting our event every 2 secondes.
func JsTestRegisterEventReactor(rc *progpAPI.SharedResourceContainer, myFunction progpAPI.JsFunction) {
	if gEventReactor != nil {
		return
	}

	// This function is now our event reactor.
	gEventReactor = myFunction

	// In a logic of event / reactor, the javascript call a function to say
	// hello I'm the function to call when the event xyz occurs.
	//
	// Since this function is used as an event reactor, it can be called more than one time.
	// But the default behaviors is to destroy the function after the first call. It's why
	// we call KeepAlive, which ask to not destroy our function. It will only be destroyed
	// by Go garbage collector if you have no more reference on the function.
	//
	// Here our function will automatically became an async function
	// which is a little slower to execute but more safe in high load scenarios.
	// And if your aren't in a such scenario you will never see a real difference.
	//
	myFunction.KeepAlive()

	// Now say that we want to automatically dispose the resource created when calling this function.
	// It's importance because without that resource are dispose through manual call to progpDispose(myResource)
	// or when the script ends.
	//
	myFunction.EnabledResourcesAutoDisposing(rc)

	// This will avoid that the VM exit.
	rc.GetScriptContext().IncreaseRefCount()

	// We create a Go-Routine, which is an optimized thread.
	// This will allow to let the javascript engine resume his execution without blocking the whole.
	//
	progpAPI.SafeGoRoutine(func() {
		// Infinite loop
		for {
			// Wait 2 secondes
			progpAPI.PauseMs(2000)
			triggerMyEvent()
		}
	})
}

// JsTestCreateResource create a simple resource.
// It will help us test the auto-disposing mechanism.
func JsTestCreateResource(rc *progpAPI.SharedResourceContainer, aSimpleString string) *progpAPI.SharedResource {
	myResource := rc.NewSharedResource(aSimpleString, func(value any) {
		// Adding a dispose function is optional.
		// This function will be called when the resource wrapper is disposed.

		println("Disposing resource ", value.(string))
	})

	// You can dispose the resource yourself by calling Exit.
	// This will unregister the wrapper and call the dispose function for the resource.
	//
	//myResource.Exit()

	return myResource
}
