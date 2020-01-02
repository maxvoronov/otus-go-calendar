package cmd

import (
	"fmt"
	"github.com/maxvoronov/otus-go-calendar/api"
	"github.com/maxvoronov/otus-go-calendar/storage/inmemory"
	"github.com/spf13/cobra"
	"time"
)

var (
	paramHost           string
	paramPort           string
	paramReadTimeout    uint8
	paramWriteTimeout   uint8
	paramMaxHeaderBytes int

	apiServerCmd = &cobra.Command{
		Use:   "api-server",
		Short: "Run REST API server",
		Run: func(cmd *cobra.Command, args []string) {
			config := &api.ServerConfig{
				Host:           paramHost,
				Port:           paramPort,
				ReadTimeout:    time.Duration(paramReadTimeout) * time.Second,
				WriteTimeout:   time.Duration(paramWriteTimeout) * time.Second,
				MaxHeaderBytes: paramMaxHeaderBytes,
			}

			storage := inmemory.NewInMemoryStorage()

			fmt.Println("Starting API server...")
			fmt.Printf("  - Options: %+v\n", config)
			api.StartServer(storage, config)
		},
	}
)

func init() {
	apiServerCmd.PersistentFlags().StringVar(&paramHost, "host", "", "Listening host")
	apiServerCmd.PersistentFlags().StringVar(&paramPort, "port", "8080", "Listening port")
	apiServerCmd.PersistentFlags().Uint8Var(&paramReadTimeout, "read-timeout", 10, "Read timeout in seconds")
	apiServerCmd.PersistentFlags().Uint8Var(&paramWriteTimeout, "write-timeout", 10, "Write timeout in seconds")
	apiServerCmd.PersistentFlags().IntVar(&paramMaxHeaderBytes, "max-header-size", 1<<20, "Max header size in bytes")
}
