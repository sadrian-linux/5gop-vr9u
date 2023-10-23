/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getWorkspaceIDCmd represents the getWorkspaceID command
var getWorkspaceIDCmd = &cobra.Command{
	Use:   "getWorkspaceID",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getWorkspaceID called")
	},
}

func init() {
	terraformCmd.AddCommand(getWorkspaceIDCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getWorkspaceIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getWorkspaceIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
