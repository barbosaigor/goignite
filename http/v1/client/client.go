package gihttp

import (
	"context"
	"net"
	"net/http"

	gilog "github.com/b2wdigital/goignite/log"
)

func NewClient(options *Options) *http.Client {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   options.Timeout,
			KeepAlive: options.KeepAlive,
			DualStack: options.DualStack,
		}).DialContext,
		TLSHandshakeTimeout:   options.TLSHandshakeTimeout,
		DisableKeepAlives:     options.DisableKeepAlives,
		MaxIdleConns:          options.MaxIdleConn,
		MaxIdleConnsPerHost:   options.MaxIdleConnPerHost,
		MaxConnsPerHost:       options.MaxConnsPerHost,
		IdleConnTimeout:       options.IdleConnTimeout,
		ForceAttemptHTTP2:     options.ForceHTTP2,
		ExpectContinueTimeout: options.ExpectContinueTimeout,
	}

	return &http.Client{Transport: tr}
}

func NewDefaultClient(ctx context.Context) *http.Client {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(o)
}
