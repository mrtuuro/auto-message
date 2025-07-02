package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mrtuuro/auto-messager/internal/application"
	"github.com/mrtuuro/auto-messager/internal/code"
	"github.com/mrtuuro/auto-messager/internal/response"
)

// AutoStart godoc
//
// @Summary  Start the automatic 2-minute sender
// @Tags     autosend
// @Success  204
// @Router   /v1/auto/start [post]
func AutoStart(app *application.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		app.Scheduler.Start()
		return response.RespondSuccess[any](c, code.SuccessOperationCompleted, nil)
	}
}

// AutoStop godoc
//
// @Summary  Stop (pause) the automatic sender
// @Tags     autosend
// @Success  204
// @Router   /v1/auto/stop [post]
func AutoStop(app *application.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		app.Scheduler.Stop()
		return response.RespondSuccess[any](c, code.SuccessOperationCompleted, nil)
	}
}
