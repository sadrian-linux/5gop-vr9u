package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func UploadConfigVersion(
	workspace_id string,
	config_version_payload []byte,
	token string) {

	url := fmt.Sprintf(
		"https://app.terraform.io/api/v2/workspaces/%s/configuration-versions",
		workspace_id)
	log.Print(url)

	client := &http.Client{}

	request, err := http.NewRequest(
		"POST", url, bytes.NewBuffer(config_version_payload))
	if err != nil {
		log.Fatalf(
			"Something went wrong while uploading the config version: %v",
			err)
	}

	auth_header := fmt.Sprintf("Bearer %s", token)

	request.Header.Add("Authorization", auth_header)
	request.Header.Add("Content-Type", "application/vnd.api+json")

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Something went wrong while performing the request: %v",
			err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
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

	var response_map map[string]interface{}
	err = json.Unmarshal(body, &response_map)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	data := response_map["data"].(map[string]interface{})
	attributes := data["attributes"].(map[string]interface{})
	uploadURL := attributes["upload-url"].(string)

	fmt.Println("Upload URL:", uploadURL)

}
