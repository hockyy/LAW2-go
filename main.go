package main

import (
	gen "LAW2-go/gen/openapi"
	"LAW2-go/handler"
	"flag"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	swagger, err := gen.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	applicationServer := handler.NewApplicationServer()

	e := echo.New()
	e.Static("/swaggerui", "./swaggerui")
	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))
	gen.RegisterHandlers(e, applicationServer)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
