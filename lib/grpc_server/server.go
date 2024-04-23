package grpc_server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GRPCConfig struct {
	Host string `yaml:"host"` // Server Host
	Port int    `yaml:"port"` // Server Port
}

type GRPCServer struct {
	Ser  *grpc.Server
	Addr string
	lis  net.Listener
}

func NewGRPCServer(cfg *GRPCConfig, opts ...grpc.ServerOption) (*GRPCServer, error) {
	c := &GRPCServer{
		Ser:  grpc.NewServer(opts...),
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	var err error
	if c.lis, err = net.Listen("tcp", c.Addr); err != nil {
		return nil, err
	}

	return c, nil
}

func (ser *GRPCServer) Run() {
	go func() {
		_ = ser.Ser.Serve(ser.lis)
	}()
}

func (ser *GRPCServer) Close() {
	if ser.lis == nil {
		return
	}
	ser.Ser.GracefulStop()
}
