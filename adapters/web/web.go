package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/captaincodeman/clean-go/engine"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// NewWebAdapter creates a new web adaptor which will
// handle the web interface and pass calls on to the
// engine to do the real work (that's why the engine
// factory is passed in - so anything that *it* needs
// is unknown to this).
// Because the web adapter ends up quite lightweight
// it easier to replace. We could use any one of the
// Go web routers / frameworks (Gin, Echo, Goji etc...)
// or just stick with the standard framework. Changing
// should be far less costly.
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
