package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mrtuuro/auto-messager/internal/application"
	"github.com/mrtuuro/auto-messager/internal/code"
	"github.com/mrtuuro/auto-messager/internal/response"
)

// HealthcheckHandler godoc
// @Summary      Liveness probe
// @Description  Returns 200 OK with a success
// @Tags         system
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.SwaggerSuccess
// @Router       /v1/healthz [get]
func HealthcheckHandler(app *application.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		return response.RespondSuccess[any](c, code.SuccessHealthCheck, nil)
	}
}
