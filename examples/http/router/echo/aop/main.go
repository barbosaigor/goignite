package main

import (
	"context"
	"log"

	"github.com/b2wdigital/goignite/examples/http/router/echo/aop/advice"
	"github.com/b2wdigital/goignite/examples/http/router/echo/aop/handler"
	c "github.com/b2wdigital/goignite/examples/http/router/echo/aop/model/config"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/http/router/echo"
	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wesovilabs/beyond/api"
)

const Endpoint = "app.endpoint.google"

func init() {
	config.Add(Endpoint, "/google", "google endpoint")
}

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(advice.NewTracingAdvice, "handler.Get(...)").
		WithBefore(advice.NewTracingAdviceWithPrefix("[beyond]"), "handler.*(...)...")
}

func main() {

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	c := c.Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	info.AppName = "google"

	instance := echo.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	instance.GET(c.App.Endpoint.Google, handler.Get)

	echo.Serve(ctx)
}
