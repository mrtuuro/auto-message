package router

import (
	"fmt"

	"github.com/mrtuuro/auto-messager/internal/application"
	"github.com/mrtuuro/auto-messager/internal/handler"
	"github.com/mrtuuro/auto-messager/internal/middleware"
)

func Register(app *application.Application) {
	app.E.GET("/v1/healthz", handler.HealthcheckHandler(app)).Name = "Healthcheck"

	protected := app.E.Group("/v1")
	protected.Use(middleware.CustomMiddleware(app))

	protected.POST("/auto/start", handler.AutoStart(app)).Name = "AutoStart"
	protected.POST("/auto/stop", handler.AutoStop(app)).Name = "AutoStop"
	protected.GET("/messages/list", handler.ListSentHandler(app)).Name = "ListSent"

}

func PrintRoutes(app *application.Application) {
	fmt.Println("=== ROUTES ===")
	routes := app.E.Routes()
	for _, r := range routes {
		fmt.Printf("%s - [%s]%s\n", r.Name, r.Method, r.Path)
	}
}
