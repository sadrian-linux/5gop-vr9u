package model

import (
	"fmt"
	"net/http"
)

// TODO: Modify the function to expect the variables bellow as arguments
func GetWorkspaceID(org string, workspace string, token string) {
	fmt.Println("GetWorkspaceID() called", org, workspace, token)
	url := fmt.Sprintf("https://app.terraform.io/api/v2/organizations/%s/workspaces/%s", org, workspace)

	// client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Something went wrong while retrieving the workspace ID:", err)
		return
	}

	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/vnd.api+json")

	fmt.Println(request)

}
