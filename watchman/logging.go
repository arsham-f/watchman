package main

import "fmt"

func HandleError(err error, label string) bool {
	if err == nil {
		return false
	}

	fmt.Printf("ERROR [%s]: %s\n", label, err)
	return true
}

func Warn(msg string) {
	fmt.Printf("WARNING %s\n", msg)
}

func Infof(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	fmt.Printf("[INFO]: %s\n", msg)
}
