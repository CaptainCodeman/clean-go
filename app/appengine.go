// +build appengine

package main

import (
	"net/http"

	"github.com/captaincodeman/clean/adapters/web"
	"github.com/captaincodeman/clean/engine"
	"github.com/captaincodeman/clean/providers/appengine"
)

func init() {
	s := appengine.NewStorage()
	e := engine.NewEngine(s)
	http.Handle("/", web.NewWebAdapter(e))
}
