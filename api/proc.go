package api

import (
	"github.com/kunit/rprocs/proc"
	"github.com/labstack/echo"
	"net/http"
)

// GetProc GET /v1/proc
func GetProc(c echo.Context) error {
	procs, err := proc.GetProc()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, procs)
}
