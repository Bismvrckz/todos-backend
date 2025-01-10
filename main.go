package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"task-api/config"
	"task-api/databases"
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

	a.Tkbai.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.WebHost, config.APIHost},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PATCH, echo.PUT, echo.POST, echo.DELETE},
	}))

	a.Tkbai.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PATCH, echo.PUT, echo.POST, echo.DELETE},
	}))
	a.Tkbai.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            2592000,
		ContentSecurityPolicy: "default-src 'self' ;font-src 'self' fonts.googleapis.com fonts.gstatic.com; style-src 'nonce-" + config.StyleSrcNonce + "' 'self' fonts.googleapis.com fonts.gstatic.com; script-src 'self' 'nonce-" + config.ScriptSrcNonce + "' ; img-src data://* 'self' www.w3.org ",
	}))

	//'nonce-" + config.StyleSrcNonce + "' 'self' fonts.googleapis.com fonts.gstatic.com

	//logging
	initLoggingMiddleware(a)

	//init handler
	//handler.InitErrHandler(a)

	//add routes
	//routes.BuildRoutes(a)

	err := databases.ConnectAppDatabase()
	if err != nil {
		config.LogErr(err, "Error connecting to database")
		log.Fatal(err)
	}

	a.Tkbai.Logger.Fatal(a.Tkbai.Start(config.SERVERPort))
}
