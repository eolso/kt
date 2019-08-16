package cmd

import (
	"fmt"

	"github.com/ericolsonnv/kt/pkg/labrador"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node [NODENAME]",
	Short: "runs kubectl top on all pods in a node",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkSortFlag(cmd); err != nil {
			fmt.Println(err.Error())
			return
		}

		if quietFlag {
			labrador.ShowProgress = false
		}

		pods := labrador.FetchNode(args[0])
		if sortFlag != "" {
			labrador.SortPods(pods, sortFlag)
		}

		labrador.PrettyPrint(pods)
	},
}

func init() {
	nodeCmd.Flags().BoolVarP(&quietFlag, "quiet", "q", false, "disable progress bar")
	nodeCmd.Flags().StringVarP(&sortFlag, "sort", "s", "", "sort output by [name|memory|cpu]")
	rootCmd.AddCommand(nodeCmd)
}
