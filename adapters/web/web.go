package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/captaincodeman/clean/engine"
)

type (
	adapter struct {
		engine.EngineFactory
	}
)

func NewWebAdapter(f engine.EngineFactory) http.Handler {
	a := &adapter{f}
	e := echo.New()

	e.Use(middleware.Logger())
	e.SetRenderer(NewTemplate())

	initGreetings(e, a, "/")

	return e
}
