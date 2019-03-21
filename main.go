package main

import (
	"github.com/kunit/rprocs/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	{
		v1.GET("/proc", api.GetProc)
	}

	e.Start(":8000")
}
