package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-terra-parse",
	Short: "Parse terraform files",
	Long:  `Parse terraform files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root Command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
