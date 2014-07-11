package main

import (
	. "gopkgs.com/magick.v1"
)

var (
	imageA    = &Image{}
	imageB    = &Image{}
	NullImage = Image{}
	images    = make(chan *Snapshot, 2)
)

type Snapshot struct {
	img  *Image
	name string
}

/*
	Blocks until there is an image available
*/
func NextSnapshot() *Snapshot {
	return <-images
}

/*
	Continuously keep the buffer full
*/
func DownloadImages() {

	for {

		nextFile := PopFromQueue()
		data, err := Get(nextFile)

		if HandleError(err, "Downloading file") {
			continue
		}

		go AddToQueue(nextFile, data)
		go Del(nextFile)

	}
}

func AddToQueue(fname string, data []byte) {

	image, err := DecodeData(data)

	if HandleError(err, "Decoding image") {
		return
	}

	images <- &Snapshot{
		img:  image,
		name: fname,
	}
}

/*
	Rotate images calculate mean difference
*/
func NextImageAndCompare() (string, float64) {
	var s Stopwatch

	*imageA = *imageB

	next := NextSnapshot()
	imageB = next.img

	if *imageA == NullImage || *imageB == NullImage {
		return "", 0
	}

	s.Start()
	diff, err := imageB.Compare(imageA)
	s.Stop("Compare")

	if HandleError(err, "Comparing two images") {
		return "", 0
	}

	return next.name, diff.MeanPerPixel
}
