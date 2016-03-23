package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/captaincodeman/clean/engine"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func NewWebAdapter(f engine.EngineFactory, log bool) http.Handler {
	var e *gin.Engine
	if log {
		e = gin.Default()
	} else {
		e = gin.New()
	}

	e.LoadHTMLGlob("templates/*")

	initGreetings(e, f, "/")

	return e
}
