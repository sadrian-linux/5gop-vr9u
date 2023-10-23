package cmd

import (
	"fmt"
	"log"
	"orb/model"

	"github.com/spf13/cobra"
)

// terraformCmd represents the terraform command
var terraformPushCodeCmd = &cobra.Command{
	Use:   "terraformPushCode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("terraformPushCode called")
		// primitive checks
		// if len(args) < 1 {
		// 	log.Fatal("No arguments provided.")
		// }
		// argument := args[0]
		// fmt.Println(argument)

		// call CreateArchive()
		model.CreateArchive()
		org, _ := cmd.Flags().GetString("org")
		workspace, _ := cmd.Flags().GetString("workspace")
		token, _ := cmd.Flags().GetString("token")

		// call GetWorkspaceID()
		// primitive checks
		// TODO: enforce complex checks
		if org != "" && workspace != "" && token != "" {
			model.GetWorkspaceID(org, workspace, token)
		} else {
			log.Fatal("GetWorkspaceID: Invalid or missing token, org or workspace.")
		}

	},
}

func init() {
	rootCmd.AddCommand(terraformPushCodeCmd)

	// flag definitions
	terraformPushCodeCmd.Flags().StringP("token", "t", "", "access token")
	terraformPushCodeCmd.Flags().StringP("org", "o", "", "organization name")
	terraformPushCodeCmd.Flags().StringP("workspace", "w", "", "workspace name")
	// terraformPushCodeCmd.PersistentFlags().String("foo", "", "A help for foo")
	// terraformPushCodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
