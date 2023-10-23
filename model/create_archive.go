package model

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CreateArchive() (string, error) {
	fmt.Println("CreateArchive() called")
	//TODO: compose the value of archive_name dynamically
	archive := "/tmp/test.tar.gz"
	cmd := exec.Command("tar", "zcf", archive, ".")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf("Error: %v\n%s", err, stderr.String())
	}

	return archive, nil
}
