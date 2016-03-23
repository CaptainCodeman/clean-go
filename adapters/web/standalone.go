// +build !appengine

package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func ctx(c *gin.Context) context.Context {
	return c
}
