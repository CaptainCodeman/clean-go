// +build appengine

package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func ctx(c *gin.Context) context.Context {
	return appengine.NewContext(c.Request)
}
