package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	"github.com/b2wdigital/goignite/pkg/transport/client/resty/v2"
)

func main() {

	var err error

	config.Load()

	ctx := context.Background()

	log.NewLogger(logrus.NewLogger())

	l := log.FromContext(ctx)

	client := resty.NewClient(ctx, &resty.Options{})
	request := client.R().EnableTrace()

	response, err := request.Get("http://google.com")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(response.String())
}