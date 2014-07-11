package main

import (
	"flag"
	"time"
)

var (
	produce = flag.Bool("produce", false, "Produce snapshots")
	consume = flag.Bool("consume", false, "Consume snapshots")
)

func main() {
	flag.Parse()
	if err := ConnectDropbox(); HandleError(err, "Connecting to dropbox") {
		return
	}

	if err := ConnectRedis(); HandleError(err, "Connecting to redis") {
		return
	}

	if *consume {
		go Consumer()
	}

	if *produce {
		go Producer()
	}

	for {
		time.Sleep(10000)
	}
}

func Consumer() {
	Infof("Starting")
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

func Producer() {
	StartCapture()
}
