package grpch

import (
	"context"
	"net"

	"github.com/hnit-acm/hfunc/serverh"
	"google.golang.org/grpc"
)

var _ serverh.Server = (*Server)(nil)

// Server is a gRPC serverh wrapper.
type Server struct {
	*grpc.Server

	network string
	addr    string
}

// NewServer creates a gRPC serverh by options.
func NewServer(network, addr string, opts ...grpc.ServerOption) *Server {
	return &Server{
		network: network,
		addr:    addr,
		Server:  grpc.NewServer(opts...),
	}
}

// Start start the gRPC serverh.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the gRPC serverh.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
