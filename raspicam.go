package main

import (
	"fmt"
	"os/exec"
)

const (
	capDirectory    = "captures"
	captureTime     = "600000" // 10 minutes
	captureInterval = "1000"
	captureName     = "%04d"
)

func StartCapture() {
	out := fmt.Sprintf("%s/%s.jpg", captureName, capDirectory)
	cmd := exec.Command("raspistill", "-t", captureTime, "-o", out, "-tl", captureInterval)
	output, err := cmd.Output()
	if HandleError(err, "Running raspistill") {
		return
	}

	Infof("Output %s", output)
}
