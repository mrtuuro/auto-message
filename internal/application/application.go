package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mrtuuro/auto-messager/internal/autosend"
	"github.com/mrtuuro/auto-messager/internal/config"
	"github.com/mrtuuro/auto-messager/internal/service"
	"github.com/mrtuuro/auto-messager/internal/token"
	"github.com/mrtuuro/auto-messager/internal/validator"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/mrtuuro/auto-messager/docs"
)

type Application struct {
	Cfg *config.Config

	E            *echo.Echo
	TokenManager *token.TokenManager

	// SERVICES
	MessageService service.MessageService

	Scheduler *autosend.Scheduler
}

func NewApp(cfg *config.Config, msgSvc service.MessageService, sched *autosend.Scheduler) *Application {
	app := &Application{}
	tm := token.NewTokenManager(cfg.SecretKey)

	app.Cfg = cfg
	app.MessageService = msgSvc
	app.Scheduler = sched
	app.E = setupEcho()
	app.TokenManager = tm

	return app
}

func setupEcho() *echo.Echo {
	e := echo.New()

	e.Validator = validator.NewCustomValidator()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.ERROR)
	return e
}

func (app *Application) Run(port string) {
	ctx, stop := signal.NotifyContext(app.Cfg.Ctx, os.Interrupt)
	defer stop()

	go func() {
		if err := app.E.Start(port); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
			app.E.Logger.Fatal("Shutting down the server!")
		}
	}()

	<-ctx.Done()
	app.Scheduler.Stop()
	ctx, cancel := context.WithTimeout(app.Cfg.Ctx, 10*time.Second)
	defer cancel()
	if err := app.E.Shutdown(ctx); err != nil {
		app.E.Logger.Fatal(err)
	}
}
