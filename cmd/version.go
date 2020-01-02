package cmd

import (
	"fmt"
	"github.com/maxvoronov/otus-go-calendar/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show information about version, commit and build time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:    ", version.Version)
		fmt.Println("Commit:     ", version.Commit)
		fmt.Println("Build time: ", version.BuildTime)
	},
}
