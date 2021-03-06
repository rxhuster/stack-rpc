package grpc

import (
	"github.com/stack-labs/stack-rpc/plugin"
	"github.com/stack-labs/stack-rpc/transport"
)

type grpcTransportPlugin struct {
}

func (g *grpcTransportPlugin) Name() string {
	return "grpc"
}

func (g *grpcTransportPlugin) Options() []transport.Option {
	return nil
}

func (g *grpcTransportPlugin) New(opts ...transport.Option) transport.Transport {
	return NewTransport(opts...)
}

func init() {
	plugin.TransportPlugins["grpc"] = &grpcTransportPlugin{}
}
