package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/captaincodeman/clean/engine"
)

type (
	greeter struct {
		engine.Greeter
	}
)

func initGreetings(e *gin.Engine, f engine.EngineFactory, endpoint string) {
	greeter := &greeter{f.NewGreeter()}
	g := e.Group(endpoint)
	{
		g.GET("", greeter.list)
		g.POST("", greeter.add)
	}
}

func (g greeter) list(c *gin.Context) {
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil || count == 0 {
		count = 5
	}
	req := &engine.ListGreetingsRequest{
		Count: count,
	}
	res := g.List(ctx(c), req)
	if c.Query("format") == "json" {
		c.JSON(http.StatusOK, res.Greetings)
	} else {
		c.HTML(http.StatusOK, "guestbook.html", res)
	}
}

func (g greeter) add(c *gin.Context) {
	req := &engine.AddGreetingRequest{
		Author:  c.PostForm("Author"),
		Content: c.PostForm("Content"),
	}
	g.Add(ctx(c), req)
	// TODO: set flash cookie for "added"
	c.Redirect(http.StatusFound, "/")
}
