package controller

import (
	"fmt"
	"orb/cmd"
	"orb/model"
	"os"

	"github.com/spf13/cobra"
)

// terraformCmd represents the terraform command
var terraformPushCodeCmd = &cobra.Command{
	Use:   "terraformPushCode",
	Short: "push terraform configuration files",
	Long:  `orb terraformPushCode -o=<org_name> -w=<workspace_name> -t=<token>`,
	Run: func(cmd *cobra.Command, args []string) {

		// exit if unexpected arguments are provided
		if len(args) != 0 {
			fmt.Println("Error: no arguments expected.")
			cmd.Usage()
			os.Exit(1)
		}

		// call CreateArchive()
		archive := model.CreateArchive()
		fmt.Printf("Archive path: %s\n", archive)

		// call GetWorkspaceID()
		org, _ := cmd.Flags().GetString("org")
		workspace, _ := cmd.Flags().GetString("workspace")
		token, _ := cmd.Flags().GetString("token")

		workspaceID := model.GetWorkspaceID(org, workspace, token)
		fmt.Printf("Workspace ID: %s\n", workspaceID)

	},
}

func init() {
	cmd.RootCmd.AddCommand(terraformPushCodeCmd)

	// flag definitions
	terraformPushCodeCmd.Flags().StringP("org", "o", "", "required: -o=<organization_name>")
	terraformPushCodeCmd.Flags().StringP("workspace", "w", "", "required: -w=<workspace_name>")
	terraformPushCodeCmd.Flags().StringP("token", "t", "", "required: -t=<access_token>")

	// required flags
	terraformPushCodeCmd.MarkFlagRequired("org")
	terraformPushCodeCmd.MarkFlagRequired("workspace")
	terraformPushCodeCmd.MarkFlagRequired("token")
}
