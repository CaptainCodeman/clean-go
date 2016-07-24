// +build appengine

package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

// getContext provides the appengine context that
// all of the appengine services rely on. Notice
// the build tag that only includes this version
// if we're running on appengine. The standalone
// version is in standalong.go
func getContext(c *gin.Context) context.Context {
	return appengine.NewContext(c.Request)
}
