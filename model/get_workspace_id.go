package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetWorkspaceID(org string, workspace string, token string) string {
	log.Println("GetWorkspaceID() called", org, workspace)

	url := fmt.Sprintf(
		"https://app.terraform.io/api/v2/organizations/%s/workspaces/%s",
		org,
		workspace)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf(
			"Something went wrong while retrieving the workspace ID: %v",
			err)
	}

	auth_header := fmt.Sprintf("Bearer %s", token)

	request.Header.Add("Authorization", auth_header)
	request.Header.Add("Content-Type", "application/vnd.api+json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error performing the request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf(
			"API request has failed with status code: %d",
			response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf(
			"An error was encountered while reading the response body: %v",
			err)
	}

	var responseMap map[string]interface{}

	if err := json.Unmarshal(body, &responseMap); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	data := responseMap["data"].(map[string]interface{})
	workspace_id := data["id"].(string)

	return workspace_id
}
