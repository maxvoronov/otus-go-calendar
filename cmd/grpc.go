package cmd

import (
	"fmt"
	"github.com/maxvoronov/otus-go-calendar/grpc"
	"github.com/spf13/cobra"
	"time"
)

var (
	grpcParamHost        string
	grpcParamPort        string
	grpcParamConnTimeout uint8

	grpcServerCmd = &cobra.Command{
		Use:   "grpc-server",
		Short: "Run gRPC server",
		Run: func(cmd *cobra.Command, args []string) {
			serverConfig := &grpc.ServerConfig{
				Storage:     appConfig.Storage,
				Logger:      appConfig.Logger,
				Port:        grpcParamPort,
				ConnTimeout: time.Duration(grpcParamConnTimeout) * time.Second,
			}

			fmt.Println("Starting gRPC server...")
			grpc.StartServer(serverConfig)
		},
	}
)

func init() {
	grpcServerCmd.PersistentFlags().StringVar(&grpcParamHost, "host", "0.0.0.0", "Listening host")
	grpcServerCmd.PersistentFlags().StringVar(&grpcParamPort, "port", "6565", "Listening port")
	grpcServerCmd.PersistentFlags().Uint8Var(&grpcParamConnTimeout, "timeout", 15, "Connection timeout in seconds")
}
