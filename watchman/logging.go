package main

import (
	. "fmt"
)

func HandleError(err error, label string) bool {
	if err == nil {
		return false
	}

	Printf("[ERROR] (%s): %s\n", label, err)
	return true
}

func Warn(msg string) {
	Printf("[WARNING] %s\n", msg)
}

func Infof(format string, params ...interface{}) {
	msg := Sprintf(format, params...)
	Printf("[INFO]: %s\n", msg)
}
