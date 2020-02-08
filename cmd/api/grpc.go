package main

import (
	"fmt"
	"os"
	"time"

	"github.com/maxvoronov/otus-go-calendar/grpc"
	"github.com/spf13/cobra"
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
				Port:        grpcParamPort,
				ConnTimeout: time.Duration(grpcParamConnTimeout) * time.Second,
			}

			fmt.Println("Starting gRPC server...")
			grpcServer, err := grpc.InitializeServer()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			grpcServer.Start(serverConfig)
		},
	}
)

func init() {
	grpcServerCmd.PersistentFlags().StringVar(&grpcParamHost, "host", "0.0.0.0", "Listening host")
	grpcServerCmd.PersistentFlags().StringVar(&grpcParamPort, "port", "6565", "Listening port")
	grpcServerCmd.PersistentFlags().Uint8Var(&grpcParamConnTimeout, "timeout", 15, "Connection timeout in seconds")
}
