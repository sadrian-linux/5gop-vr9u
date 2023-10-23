/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"orb/model"

	"github.com/spf13/cobra"
)

// createArchiveCmd represents the createArchive command
var createArchiveCmd = &cobra.Command{
	Use:   "createArchive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createArchive called")
		model.CreateArchive()
	},
}

func init() {
	terraformCmd.AddCommand(createArchiveCmd)

	// flag definitions
	// createArchiveCmd.PersistentFlags().BoolP("help", "h", false, helpVar)
	createArchiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
