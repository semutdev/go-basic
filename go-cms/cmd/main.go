package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/semutdev/go-cms/internal/handler"
	"github.com/semutdev/go-cms/views/pages"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	// e.GET("/", func(c echo.Context) error {
	// 	return handler.Render(c, pages.Home())
	// })

	e.GET("/", func(c echo.Context) error {
		return handler.Render(c, pages.Dashboard())
	})

	e.Logger.Fatal(e.Start("localhost:8181"))
}