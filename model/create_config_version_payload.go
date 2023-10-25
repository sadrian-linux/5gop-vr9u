package model

import (
	"encoding/json"
	"log"
)

func CreateConfigVersionPayload() []byte {
	spec_mode := true

	type Attributes struct {
		Queue       bool `json:"auto-queue-runs"`
		Speculative bool `json:"speculative"`
	}

	type Data struct {
		Type       string     `json:"type"`
		Attributes Attributes `json:"attributes"`
	}

	json_struc := Data{
		Type:       "configuration-versions",
		Attributes: Attributes{false, spec_mode},
	}

	config_version_file := map[string]interface{}{"data": json_struc}

	bytes, err := json.Marshal(config_version_file)
	if err != nil {
		log.Fatalf(
			"Something went wrong while encoding data in JSON format: %v",
			err)
	}

	return bytes
}
