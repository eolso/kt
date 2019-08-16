package cmd

import (
	"fmt"
	"kt/pkg/labrador"

	"github.com/spf13/cobra"
)

var sortFlag string

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node [NODENAME]",
	Short: "runs kubectl top on all pods in a node",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("sort") {
			if err := checkSortFlag(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		pods := labrador.FetchNode(args[0])
		labrador.PrettyPrint(pods)
	},
}

func checkSortFlag() (err error) {
	switch sortFlag {
	case "name":
		return nil
	case "memory":
		return nil
	case "cpu":
		return nil
	case "":
		return nil
	default:
		return fmt.Errorf("Error: sort specified does not match [name|memory|cpu]")
	}
}

func init() {
	nodeCmd.Flags().StringVarP(&sortFlag, "sort", "s", "", "sort output by [name|memory|cpu]")
	rootCmd.AddCommand(nodeCmd)
}
