package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)
	e.GET("/", getHelloWorld)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(ctx echo.Context) error {
	// User ID from path `users/:id`
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func getHelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World!")
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
