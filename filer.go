// Copyright (c) 2014 Ben Morgan <neembi@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

var (
	listen   = flag.String("listen", "localhost:34537", "listen on this address")
	username = flag.String("user", "nobody", "username for authentication")
	password = flag.String("pass", "", "password for authentication")
)

func main() {
	flag.Parse()

	root := "."
	if flag.NArg() == 1 {
		root = flag.Arg(0)
	}
	log.Printf("Serving filesystem at %s", root)

	var handler http.Handler
	if *username != "" && *password != "" {
		log.Printf("Requiring authentication as %s", *username)
		handler = httpauth.SimpleBasicAuth(*username, *password)(http.FileServer(http.Dir(root)))
	} else {
		log.Println("Warning: no authentication required")
		handler = http.FileServer(http.Dir(root))
	}

	log.Printf("Listening on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, handler))
}
