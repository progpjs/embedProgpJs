function r(t, v) { progpPrint("Function '", t, "' result is: ", v); }

function doTestGetReturnValue() {
    r("testByteArray", progpBufferToString(testByteArray(progpStringToBuffer("my text"))))
    r("testStructPtr", testStructPtr({Name: "NA", Forename: "FNA"}))
    r("testStructRef", testStructRef({Name: "NA", Forename: "FNA"}))
    r("testStringArray", testStringArray(["aa", "bb"]))
    r("testInt", testInt(1234))
    r("testInt32", testInt32(1234))
    r("testInt64", testInt64(1234))
    r("testFloat32", testFloat32(12.34))
    r("testFloat64", testFloat64(12.34))
    r("testString", testString("myString"))
    r("testBool - true", testBool(true))
    r("testBool - false", testBool(false))
}

function catchError(who, f) {
    let hasError = false;

    try {
        f()
    } catch (e) {
        hasError = true;
        progpPrint(who, " throws ", e.toString())
    }

    if (!hasError) {
        progpPrint(who, " don't throw error")
    }
}

function doTestCallJavascript() {
    testCallbackWithString((err, res) => { r("testCallbackWithString", res) }, "stringValue")
    testCallbackWithBool((err, res) => { r("testCallbackWithBool(true)", res) }, true)
    testCallbackWithBool((err, res) => { r("testCallbackWithBool(false)", res) }, false)
    testCallbackWithDouble((err, res) => { r("testCallbackWithDouble", res) }, 5.678)
    testCallbackWithError((err) => { r("testCallbackWithError", err) }, "my error message")
    testCallbackWithoutValues((err, res) => { r("testCallbackWithoutValues", res) })
    testCallbackWithStringBuffer((err, res) => { r("testCallbackWithStringBuffer", res) }, "stringValue")
    testCallbackWithArrayBuffer((err, res) => { r("testCallbackWithArrayBuffer", progpBufferToString(res)) }, progpStringToBuffer("myArrayBuffer"))
}

function doTestCallAsync() {
    testAsync("my name", (err, res)=> { r("testAsync", res) });
}

function doTestThrowError() {
    catchError("testThrowError", () => testThrowError("boom"));
    catchError("testThrowErrorP1", () => testThrowErrorP1("boom"));
    catchError("testThrowErrorP2", () => testThrowErrorP2("boom"));
}

function doTests() {
    doTestGetReturnValue();
    doTestCallJavascript();
    doTestCallAsync();
    doTestThrowError();
}

doTests();