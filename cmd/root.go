/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "axa-notebook-ssh",
	Short: "A proxy to access the AXA Notebook SSH server",
	Long:  `A proxy to access the AXA Notebook SSH server. It uses Cloudflare tunnel to access the AXA Notebook SSH server.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
