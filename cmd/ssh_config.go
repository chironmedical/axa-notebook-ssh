/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sshConfigCmd represents the sshConfig command
var sshConfigCmd = &cobra.Command{
	Use:   "ssh-config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(*cobra.Command, []string) {
		fmt.Println("sshConfig called")

	},
}

func init() {
	rootCmd.AddCommand(sshConfigCmd)
}
