package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/engine"
)

type (
	greeter func(context.Context) *engine.Greeter
)

func initGreetings(e *echo.Echo, a *adapter, endpoint string) {
	g := e.Group(endpoint)
	g.Get("", a.Greeter().List)
	g.Post("", a.Greeter().Add)
}

func (a *adapter) Greeter() greeter {
	return a.GetGreeter
}

func (g greeter) List(c *echo.Context) error {
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil || count == 0 {
		count = 5
	}
	req := &engine.ListGreetingsRequest{
		Count: count,
	}
	res := g(c).List(req)
	if c.Query("format") == "json" {
		return c.JSONIndent(http.StatusOK, res.Greetings, "", "  ")
	} else {
		return c.Render(http.StatusOK, "guestbook.html", res)
	}
}

func (g greeter) Add(c *echo.Context) error {
	req := &engine.AddGreetingRequest{
		Author:  c.Form("Author"),
		Content: c.Form("Content"),
	}
	g(c).Add(req)
	// TODO: set flash cookie
	return c.Redirect(http.StatusFound, "/")
}
