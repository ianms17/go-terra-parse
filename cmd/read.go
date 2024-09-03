package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a terraform project and output to a file",
	Long:  `Read through a terraform project and extract metadata, and output that to a file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Read Command")
	},
}
