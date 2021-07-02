package hhttp

import (
	"context"
	"crypto/tls"
	"github.com/hnit-acm/hfunc/hutils"
	"net"
	"net/http"
	"time"
)

// Server is a HTTP hserver wrapper.
type Server struct {
	*http.Server

	network string
	addr    string
	opts    options
}

// NewServer creates a HTTP hserver by options.
func NewServer(network, addr string, opts ...Option) *Server {
	options := options{
		readTimeout:  time.Second,
		writeTimeout: time.Second,
		idleTimeout:  time.Minute,
	}
	for _, o := range opts {
		o.apply(&options)
	}
	return &Server{
		network: network,
		addr:    addr,
		opts:    options,
		Server: &http.Server{
			Handler:      options.handler,
			TLSConfig:    options.tlsConfig,
			ReadTimeout:  options.readTimeout,
			WriteTimeout: options.writeTimeout,
			IdleTimeout:  options.idleTimeout,
		},
	}
}

// Start start the HTTP hserver.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the HTTP hserver.
func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}

func New(ops ...Option) *http.Server {
	op := options{
		readTimeout:  time.Second,
		writeTimeout: time.Second,
		addr:         ":9999",
		idleTimeout:  time.Minute,
		handler:      http.DefaultServeMux,
		tlsConfig:    hutils.GenTLSConfigNoErr(),
	}
	for _, o := range ops {
		o.apply(&op)
	}
	return &http.Server{
		Addr:         op.addr,
		Handler:      op.handler,
		TLSConfig:    op.tlsConfig,
		ReadTimeout:  op.readTimeout,
		WriteTimeout: op.writeTimeout,
		IdleTimeout:  op.idleTimeout,
	}
}

func ListenTLS(sever *http.Server) (err error) {
	tcpConn, err := net.Listen("tcp", sever.Addr)
	if err != nil {
		return
	}
	tlsConn := tls.NewListener(tcpConn, sever.TLSConfig)
	return sever.Serve(tlsConn)
}
