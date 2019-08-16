package cmd

import (
	"github.com/ericolsonnv/kt/pkg/labrador"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod [PODNAME]",
	Short: "runs kubectl top on a specific pod",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if pod, err := labrador.FetchPod(args[0]); err != nil {
			log.Fatal(err)
		} else {
			pod.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(podCmd)
}
