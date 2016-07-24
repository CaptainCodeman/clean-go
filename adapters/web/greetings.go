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

// wire up the greetings routes
func initGreetings(e *gin.Engine, f engine.EngineFactory, endpoint string) {
	greeter := &greeter{f.NewGreeter()}
	g := e.Group(endpoint)
	{
		g.GET("", greeter.list)
		g.POST("", greeter.add)
	}
}

// list converts the parameters into an engine
// request and then marshalls the results based
// on the format requested, returning either an
// html rendered page or JSON (to simulate basic
// content negotiation). It's simpler if the UI
// is a SPA and the web interface is just an API.
func (g greeter) list(c *gin.Context) {
	ctx := getContext(c)
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil || count == 0 {
		count = 5
	}
	req := &engine.ListGreetingsRequest{
		Count: count,
	}
	res := g.List(ctx, req)
	if c.Query("format") == "json" {
		c.JSON(http.StatusOK, res.Greetings)
	} else {
		c.HTML(http.StatusOK, "guestbook.html", res)
	}
}

// add accepts a form post and creates a new
// greoting in the system. It could be made a
// lot smarter and automatically check for the
// content type to handle forms, JSON etc...
func (g greeter) add(c *gin.Context) {
	ctx := getContext(c)
	req := &engine.AddGreetingRequest{
		Author:  c.PostForm("Author"),
		Content: c.PostForm("Content"),
	}
	g.Add(ctx, req)
	// TODO: set a flash cookie for "added"
	// if this was a web request, otherwise
	// send a nice JSON response ...
	c.Redirect(http.StatusFound, "/")
}
