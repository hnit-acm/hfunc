package hhttp

import (
	"crypto/tls"
	"net/http"
	"time"
)

type (
	options struct {
		network      string
		addr         string
		handler      http.Handler
		tlsConfig    *tls.Config
		readTimeout  time.Duration
		writeTimeout time.Duration
		idleTimeout  time.Duration
	}

	Option interface {
		apply(*options)
	}

	optionFunc func(*options)
)

func (of optionFunc) apply(ops *options) {
	of(ops)
}

func WithHandler(handler http.Handler) Option {
	return optionFunc(func(o *options) {
		o.handler = handler
	})
}

// WithTLSConfig with hserver tls config.
func WithTLSConfig(conf *tls.Config) Option {
	return optionFunc(func(o *options) {
		o.tlsConfig = conf
	})
}

// WithReadTimeout with read timeout.
func WithReadTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.readTimeout = timeout
	})
}

// WithWriteTimeout with write timeout.
func WithWriteTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.writeTimeout = timeout
	})
}

// WithIdleTimeout with read timeout.
func WithIdleTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.idleTimeout = timeout
	})
}
