// +build !appengine

package main

import (
	"net/http"

	"github.com/captaincodeman/clean/adapters/web"
	"github.com/captaincodeman/clean/engine"
	"github.com/captaincodeman/clean/providers/mongodb"
)

func main() {
	s := mongodb.NewStorage(config.MongoURL)
	e := engine.NewEngine(s)
	http.ListenAndServe(":8080", web.NewWebAdapter(e, true))
}
