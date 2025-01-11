package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"strings"
	"task-api/config"
	"task-api/databases"
	"task-api/handler"
	"task-api/routes"
)

func main() {
	a := new(config.Apps)

	a.Tkbai = echo.New()

	a.Tkbai.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogErrorFunc: func(ctx echo.Context, err error, stack []byte) error {
			fmt.Println(string(stack))
			config.Log.Error().Str("REQUEST", ctx.Request().URL.Path).Msg("[PANIC]")
			return err
		},
	}))

	//'nonce-" + config.StyleSrcNonce + "' 'self' fonts.googleapis.com fonts.gstatic.com

	//logging
	initLoggingMiddleware(a)

	//init handler
	handler.InitErrHandler(a)

	//add routes
	routes.BuildRoutes(a)

	err := databases.ConnectAppDatabase()
	if err != nil {
		config.LogErr(err, "Error connecting to database")
		log.Fatal(err)
	}

	a.Tkbai.Logger.Fatal(a.Tkbai.Start(config.SERVERPort))
}

func initLoggingMiddleware(ein *config.Apps) {
	logger := config.Log

	ein.Tkbai.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogURIPath:   true,
		LogStatus:    true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogRequestID: true,
		LogError:     true,
		LogMethod:    true,
		LogValuesFunc: func(ctx echo.Context, values middleware.RequestLoggerValues) (err error) {
			if strings.Contains(values.URI, "public") {
				return err
			}

			if values.Error != nil {
				logger.Error().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("HOST", values.Host).
					Str("RequestID", values.RequestID).
					Stack().Err(values.Error).Msg("")
			} else {
				logger.Info().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("HOST", values.Host).
					Str("RequestID", values.RequestID).
					Msg("Request")
			}
			return nil
		},
	}))
}
