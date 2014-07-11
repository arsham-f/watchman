package main

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
	var name string
	var diff float64
	for {
		name, diff = NextImageAndCompare()
		Infof("Diff: %d", diff)
		if diff < 2500 {
			go Del(name)
		}
	}
}
