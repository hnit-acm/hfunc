package hgrpc

import (
	"context"
	"net"

	"github.com/hnit-acm/hfunc/hserver"
	"google.golang.org/grpc"
)

var _ hserver.Server = (*Server)(nil)

// Server is a gRPC hserver wrapper.
type Server struct {
	*grpc.Server

	network string
	addr    string
}

// NewServer creates a gRPC hserver by options.
func NewServer(network, addr string, opts ...grpc.ServerOption) *Server {
	return &Server{
		network: network,
		addr:    addr,
		Server:  grpc.NewServer(opts...),
	}
}

// Start start the gRPC hserver.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the gRPC hserver.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
