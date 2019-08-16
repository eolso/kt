package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kt",
	Short: "Wrapper script for kubectl top",
	Long: `Extends the functionality of kubectl top by allowing for entire
nodes to be topped as well. Can also be ran against single pods an even
all nodes at once.`,
}

// Execute : runs errything
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
