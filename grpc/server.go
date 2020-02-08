package grpc

import (
	"net"
	"time"

	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/sirupsen/logrus"

	eventproto "github.com/maxvoronov/otus-go-calendar/grpc/proto"
	"google.golang.org/grpc"
)

// ServerConfig contains general options for gRPC server
type ServerConfig struct {
	Host        string
	Port        string
	ConnTimeout time.Duration
}

type server struct {
	Storage domain.StorageInterface
	Logger  *logrus.Logger
}

func newServer(storage domain.StorageInterface, logger *logrus.Logger) *server {
	return &server{Storage: storage, Logger: logger}
}

// Start gRPC server
func (serv *server) Start(cfg *ServerConfig) {
	conn, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		serv.Logger.Fatalln(err)
	}

	server := grpc.NewServer(grpc.ConnectionTimeout(cfg.ConnTimeout))
	eventproto.RegisterEventServiceServer(server, serv)
	if err := server.Serve(conn); err != nil {
		serv.Logger.Fatalln(err)
	}
}
