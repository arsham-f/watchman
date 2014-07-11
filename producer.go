package main

func main() {
	if err := ConnectDropbox(); HandleError(err, "Connecting to dropbox") {
		return
	}

	if err := ConnectRedis(); HandleError(err, "Connecting to redis") {
		return
	}

	go StartCapture()
}
