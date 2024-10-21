package rpcx

import (
	"context"
	"net/url"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/smallnest/rpcx/server"
	rpcx "github.com/smallnest/rpcx/server"
)

type Server struct {
	*rpcx.Server
	baseCtx context.Context

	endpoint *url.URL
	network  string
	address  string
}

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

func NewServer(network, addr string, options ...rpcx.OptionFn) *Server {
	srv := &Server{}
	srv.init(options...)
	return srv
}

func (s *Server) init(options ...rpcx.OptionFn) {
	s.Server = server.NewServer(options...)
}

// Start starts the rpcx server.
func (s *Server) Start(ctx context.Context) error {
	log.Infof("[rpcx] server listening on: %s", s.address)
	return s.Server.Serve(s.network, s.address)
}

// Stop Stop stops the rpcx server.
func (s *Server) Stop(ctx context.Context) error {
	log.Info("[rpcx] server stopping")
	return s.Server.Close()
}

// Endpoint return a real address to registry endpoint.
// examples:
//
//	tcp@127.0.0.1:9000?isSecure=false
func (s *Server) Endpoint() (*url.URL, error) {
	addr := s.Server.Address()
	s.endpoint = &url.URL{Scheme: addr.Network(), Host: addr.String()}
	return s.endpoint, nil
}
