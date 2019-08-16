package cmd

import (
	"fmt"
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
		if err := checkSortFlag(cmd); err != nil {
			fmt.Println(err.Error())
			return
		}

		if quietFlag {
			labrador.ShowProgress = false
		}

		pods, _ := labrador.FetchPods()

		if sortFlag != "" {
			labrador.SortPods(pods, sortFlag)
		}

		labrador.PrettyPrint(pods)
	},
}

func init() {
	allCmd.Flags().BoolVarP(&quietFlag, "quiet", "q", false, "disable progress bar")
	allCmd.Flags().StringVarP(&sortFlag, "sort", "s", "", "sort output by [name|memory|cpu]")
	rootCmd.AddCommand(allCmd)
}
