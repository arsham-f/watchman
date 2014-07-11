package main

import "fmt"

func main() {
	if err := ConnectDropbox(); HandleError(err, "Connecting to dropbox") {
		return
	}

	if err := ConnectRedis(); HandleError(err, "Connecting to redis") {
		return
	}

	//Keep network and CPU operations in separate threads
	go DownloadImages()

	//Main loop
	var diff float64
	for {
		diff = NextImageAndCompare()
		fmt.Printf("Diff: %f\n", diff)
	}

}
