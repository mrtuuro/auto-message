package handler

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mrtuuro/auto-messager/internal/application"
	"github.com/mrtuuro/auto-messager/internal/code"
	"github.com/mrtuuro/auto-messager/internal/model"
	"github.com/mrtuuro/auto-messager/internal/response"
)

func ListSentHandler(app *application.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithoutCancel(c.Request().Context())
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		if limit <= 0 {
			limit = 20
		}

		msgs, err := app.MessageService.ListSent(ctx, limit, offset)
		if err != nil {
			return response.RespondError[any](c, err)
		}
		return response.RespondSuccess[[]model.Message](c, code.SuccessOperationCompleted, &msgs)

	}
}
