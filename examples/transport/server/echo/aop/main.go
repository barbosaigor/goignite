package main

import (
	"context"

	"github.com/b2wdigital/goignite/examples/transport/server/echo/aop/advice"
	"github.com/b2wdigital/goignite/examples/transport/server/echo/aop/handler"
	c "github.com/b2wdigital/goignite/examples/transport/server/echo/aop/model/config"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	"github.com/b2wdigital/goignite/pkg/transport/server/echo/v4"
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

	var err error

	config.Load()

	c := c.Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	log.NewLogger(logrus.NewLogger())

	info.AppName = "google"

	instance := echo.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	instance.GET(c.App.Endpoint.Google, handler.Get)

	echo.Serve(ctx)
}