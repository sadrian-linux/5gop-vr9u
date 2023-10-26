package controller

import (
	"fmt"
	"log"
	"orb/cmd"
	"orb/model"

	"github.com/spf13/cobra"
)

// terraformCmd represents the terraform command
var terraformPushCodeCmd = &cobra.Command{
	Use:   "terraformPushCode",
	Short: "push terraform configuration files",
	Long:  `orb terraformPushCode -o=<org_name> -w=<workspace_name> -t=<token>`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		// call CreateArchive()
		archive := model.CreateArchive()
		fmt.Printf("Archive path: %s\n", archive)

		// call GetWorkspaceID()
		org, _ := cmd.Flags().GetString("org")
		workspace, _ := cmd.Flags().GetString("workspace")
		token, _ := cmd.Flags().GetString("token")

		workspace_id := model.GetWorkspaceID(org, workspace, token)
		log.Printf("Workspace ID: %s", workspace_id)

		json_payload := model.CreateConfigVersionPayload()
		model.UploadConfigVersion(workspace_id, json_payload, token)
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
