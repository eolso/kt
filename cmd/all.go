package cmd

import (
	"kt/pkg/labrador"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "runs kubectl top on all pods across all nodes",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pods, _ := labrador.FetchPods()
		labrador.PrettyPrint(pods)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
