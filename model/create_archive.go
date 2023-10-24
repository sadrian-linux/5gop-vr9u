package model

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func CreateArchive() string {
	fmt.Println("CreateArchive() called")
	//TODO: compose the value of archive_name dynamically
	archive := "/tmp/test.tar.gz"
	cmd := exec.Command("tar", "zcf", archive, ".")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("error: %v\n%s", err, stderr.String())
	}

	return archive
}
