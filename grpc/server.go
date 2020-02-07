package grpc

import (
	"net"

	eventproto "github.com/maxvoronov/otus-go-calendar/grpc/proto"
	"google.golang.org/grpc"
)

// StartServer Start gRPC server
func StartServer(cfg *ServerConfig) {
	conn, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		cfg.Logger.Fatalln(err)
	}

	server := grpc.NewServer(grpc.ConnectionTimeout(cfg.ConnTimeout))
	eventproto.RegisterEventServiceServer(server, cfg)
	if err := server.Serve(conn); err != nil {
		cfg.Logger.Fatalln(err)
	}
}
