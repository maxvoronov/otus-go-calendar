package cmd

import (
	"fmt"
	"os"

	"github.com/maxvoronov/otus-go-calendar/internal/config"
	"github.com/spf13/cobra"
)

var (
	appConfig *config.Config
	rootCmd   = &cobra.Command{
		Use: "go-calendar [command]",
	}
)

// Execute Apply CLI commands
func Execute(cfg *config.Config) {
	appConfig = cfg
	rootCmd.AddCommand(apiServerCmd)
	rootCmd.AddCommand(grpcServerCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
