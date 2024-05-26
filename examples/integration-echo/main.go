package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", HomeHandler)
	app.Logger.Fatal(app.Start(":4000"))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's t1.Component.Render().
func Render(ctx echo.Context, statusCode int, t t1.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, Home())
}
