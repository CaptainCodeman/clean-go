// +build !appengine

package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// getContext doesn't seem to do much but this is
// the standalone version and here so we can swap
// implementations using the build tag at the top.
// Look at the appengine.go file for the appengine
// specific implementation.
func getContext(c *gin.Context) context.Context {
	return c
}
