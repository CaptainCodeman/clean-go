// +build appengine

package main

import (
	"net/http"

	"github.com/captaincodeman/clean-go/adapters/web"
	"github.com/captaincodeman/clean-go/engine"
	"github.com/captaincodeman/clean-go/providers/appengine"
)

// for appengine we don't use main to start the server
// because that is done for us by the platform. Instead
// we attach to the standard mux router. Note that we're
// using the appengine provider for storage and wiring
// it up to the engine and then the engine to the web. 
func init() {
	s := appengine.NewStorage()
	e := engine.NewEngine(s)
	http.Handle("/", web.NewWebAdapter(e, false))
}
