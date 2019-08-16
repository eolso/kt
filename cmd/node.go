package cmd

import (
	"kt/pkg/labrador"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node [NODENAME]",
	Short: "runs kubectl top on all pods in a node",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pods := labrador.FetchNode(args[0])
		labrador.PrettyPrint(pods)
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
