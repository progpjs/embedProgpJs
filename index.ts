// Here is a sample server.

// This line can be removed since progjs automatically include this line.
import "@progp/core"

// All node.js library and modules starting by "@progp" are embedded inside the executable.
// It's why you can use them without requiring a node_modules installation.
//
// The source code of this module can be found in directory modCore/embed/jsMods/@progp
// of the project modules. See: https://github.com/progpjs/modules/tree/main/modCore/embed/jsMods/%40progp/core
// You can create a directory node_modules/@progp and copy/past the file inside in order to enable code completion.
//
import {HttpServer} from "@progp/http"

// Create a server.
let server = new HttpServer(8000);

// And a hostname for this server.
// Here it will only be responding to http://localhost:8000 and not http://127.0.0.1:8000.
//
let host = server.getHost("localhost");

// Add a function call whe consulting the home page.
host.GET("/", async req => {
    //console.log("Request IP:" + req.requestIP());
    //console.log("Request path: ", req.requestPath());
    req.returnHtml(200, "Hello world")
});

// Start our server and print a message.
server.start();
console.log("Server started at http://localhost:8000")

// Warning: if you want to benchmark the server, then enable the compiled mode (see the README.md).
// If not the compiled mode, then it's the dynamic mode which is slow.