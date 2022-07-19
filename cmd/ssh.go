/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/chironmedical/axa-notebook-ssh/internal/app"
	"github.com/spf13/cobra"
)

var hostname string

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(*cobra.Command, []string) {
		err := app.Ssh(&app.Config{Hostname: hostname})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.Flags().StringVar(&hostname, "hostname", "axa-notebook.chironmedical.cloud", "Hostname of the AXA Notebook SSH server")
}
