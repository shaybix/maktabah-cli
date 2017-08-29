package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root command representing when maktabah cli is run
var RootCmd = &cobra.Command{
	Use:   "maktabah",
	Short: "Maktabah is a command line tool for Shamela Books",
	Long: `A robust tool that allows you to search, download, and manage shamela books. 

For more information see https://github.com/shaybix/maktabah-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(cmd.UsageString())
	},
}

// Execute executes the rootCmd
func Execute() {

	if c, err := RootCmd.ExecuteC(); err != nil {
		c.Println(c.UsageString())
	}
}
