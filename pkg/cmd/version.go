package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xxxVitoxxx/eshop/internal/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "version",
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintInfo()
	},
}
