package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`Redis Cluster Populator
Version: v0.0.1
Go: %v	
`, runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
